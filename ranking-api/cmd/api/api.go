package api

import (
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/api/protocols"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type API struct {
	database db.MongoDatabase
}

type UserPostParam struct {
	Name    string `json:"name"`
	Current int    `json:"current"`
	Phone   string `json:"phone"`
}

type UserUpdateCurrent struct {
	Current int    `json:"current"`
	Phone   string `json:"phone"`
}

func NewAPI(db db.MongoDatabase) *API {
	return &API{
		database: db,
	}
}

func (api *API) Create(c *gin.Context) {
	var user UserPostParam
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON"})
		return
	}

	if _, err := api.database.GetByPhone(user.Phone); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Usuario ja cadastrado",
		})
		return
	}

	userToCreate := protocols.UserPostParam{
		Phone:   user.Phone,
		Name:    user.Name,
		Current: user.Current,
	}

	if err := api.database.Create(userToCreate); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "erro interno",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (api *API) Update(c *gin.Context) {
	var user UserUpdateCurrent
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON"})
		return
	}

	userRanking, err := api.database.GetByPhone(user.Phone)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "usuario nao existe",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "usuario ja cadastrado",
			})
		}
		return
	}

	userToUpdate := protocols.UserPostParam{
		Phone:   userRanking.Phone,
		Current: user.Current,
		Name:    userRanking.Name,
	}
	if err := api.database.Update(userRanking.Id, userToUpdate); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "erro interno",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (api *API) Clean(c *gin.Context) {
	if err := api.database.DeleteAll(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "erro interno",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (api *API) Ranking(c *gin.Context) {
	users, err := api.database.GetAll()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "erro interno",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
