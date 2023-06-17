package pkg

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoClient struct {
	uri string
}

func NewMongoClient(uri string) *MongoClient {
	return &MongoClient{
		uri: uri,
	}
}

func (m *MongoClient) Client() (*mongo.Client, error) {
	return mongo.NewClient(options.Client().SetAuth(
		options.Credential{
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		}).ApplyURI(m.uri))
}
