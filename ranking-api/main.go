package main

import (
	"context"
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/api"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/pkg"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
)

func main() {
	//err := dotenv.Load()
	//if err != nil {
	//	panic("Falha ao carregar o arquivo .env")
	//}
	m := pkg.NewMongoClient(os.Getenv("DB_HOST"))
	ctxTodo := context.TODO()
	client, err := m.Client()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctxTodo)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctxTodo)
	err = client.Ping(ctxTodo, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	app := db.NewDatabase(client, "test")
	if err != nil {
		panic("Falha ao inicializar o aplicativo")
	}

	router := gin.Default()
	api := api.NewAPI(*app)

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Insira a origem permitida do seu aplicativo front-end
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	router.POST("/api/create", api.Create)
	router.PUT("/api/update", api.Update)
	router.DELETE("/api/delete", api.Clean)
	router.GET("/api/ranking", api.Ranking)

	fmt.Println("API Starting")
	log.Fatal(http.ListenAndServe(":4546", router))
}
