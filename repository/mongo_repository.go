package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoRepository interface {
	FindAllSessions() ([]*Session, error)
	InsertSession(session *Session) error
	FindById(id string) (*Session, error)
}

type mongoRepository struct {
	mongoClient *mongo.Client
}

var mongoClient *mongo.Client

func New() *mongoRepository {
	var err error
	// Connect to mongo
	mongoClient, err = connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	return &mongoRepository{
		mongoClient: mongoClient,
	}
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "mongodb",
		Password: "mongodbpw",
	})

	// connect to mongodb
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	log.Println("Connected to mongo DB")

	return c, nil
}

func DeferDisconnect() {
	// Create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Close mongo connection
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
		log.Println("Disconnected Mongo")
	}()
}
