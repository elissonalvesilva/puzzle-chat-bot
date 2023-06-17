package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/api/protocols"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	dbName string
	db     *mongo.Client
}

type UserModel struct {
	Id      string `bson:"_id" json:"_id,omitempty"`
	Name    string `bson:"name" json:"name"`
	Current int    `bson:"current" json:"current"`
	Phone   string `bson:"phone" json:"phone"`
}

func NewDatabase(db *mongo.Client, dbName string) *MongoDatabase {
	return &MongoDatabase{
		dbName: dbName,
		db:     db,
	}
}

func (d *MongoDatabase) Create(user protocols.UserPostParam) error {
	_, err := d.db.Database(d.dbName).Collection("puzzle_user").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (d *MongoDatabase) GetByPhone(phone string) (protocols.UserGetResponse, error) {
	var user protocols.UserGetResponse
	filter := bson.M{"phone": phone}
	err := d.db.Database(d.dbName).Collection("puzzle_user").FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return protocols.UserGetResponse{}, err
	}

	return user, nil
}

func (d *MongoDatabase) Update(id string, user protocols.UserPostParam) error {
	idToFind, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", idToFind}}
	update := bson.D{
		{"$set", bson.D{{"name", user.Name}}},
		{"$set", bson.D{{"phone", user.Phone}}},
		{"$set", bson.D{{"current", user.Current}}},
	}

	resp, err := d.db.Database(d.dbName).Collection("puzzle_user").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if resp.ModifiedCount <= 0 {
		err = errors.New("user not modified")
		return err
	}

	return nil
}

func (d *MongoDatabase) GetAll() (protocols.UsersRanking, error) {
	options := options.Find().SetSort(bson.D{{Key: "current", Value: 1}})

	cursor, err := d.db.Database(d.dbName).Collection("puzzle_user").Find(context.Background(), bson.D{}, options)
	if err != nil {
		return protocols.UsersRanking{}, err
	}

	var users []UserModel
	err = cursor.All(context.Background(), &users)
	if err != nil {
		fmt.Println("Erro ao decodificar os usuários:", err)
		return protocols.UsersRanking{}, err
	}

	userResponse := protocols.UsersRanking{
		Users: users,
	}

	return userResponse, nil
}

func (d *MongoDatabase) DeleteAll() error {
	_, err := d.db.Database(d.dbName).Collection("puzzle_user").DeleteMany(context.Background(), nil)
	if err != nil {
		return err
	}

	return nil
}
