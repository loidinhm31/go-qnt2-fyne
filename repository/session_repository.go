package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Session struct {
	ID          string `bson:"_id,omitempty" json:"id,omitempty"`
	SessionName string `bson:"session_name" json:"session_name"`
	SessionKey  string `bson:"session_key" json:"session_key"`
}

const (
	mongoURL        = "mongodb://mongodb:mongodbpw@localhost:27017/logs?authSource=admin"
	mongoDB         = "qnt2"
	mongoCollection = "sessions"
)

// FindAllSessions returns all session
func (mongo *mongoRepository) FindAllSessions() ([]*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := mongo.mongoClient.Database(mongoDB).Collection(mongoCollection)

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessions []*Session

	for cursor.Next(ctx) {
		var item Session

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding log into slice:", err)
			return nil, err
		} else {
			sessions = append(sessions, &item)
		}
	}
	return sessions, nil
}
