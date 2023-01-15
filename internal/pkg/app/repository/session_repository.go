package repository

import (
	"context"
	"go-qn2management/internal/pkg/app/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type Session struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	SessionName string    `bson:"session_name" json:"session_name"`
	SessionKey  string    `bson:"session_key" json:"session_key"`
	Order       int32     `bson:"order" json:"order"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

// FindAllSessions returns all session
func (mongo *mongoRepository) FindAllSessions() ([]*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoSessionsCollection)

	cursor, err := collection.Find(context.TODO(), bson.D{})
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
			log.Println("Error decoding session into slice:", err)
			return nil, err
		} else {
			sessions = append(sessions, &item)
		}
	}
	return sessions, nil
}

func (mongo *mongoRepository) InsertSession(session *Session) error {
	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoSessionsCollection)

	_, err := collection.InsertOne(context.TODO(), Session{
		SessionName: session.SessionName,
		SessionKey:  session.SessionKey,
		Order:       session.Order,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		log.Println("Error inserting into sessions", err)
		return err
	}
	return nil
}

func (mongo *mongoRepository) FindSessionById(id string) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoSessionsCollection)

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var session Session
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (mongo *mongoRepository) DeleteSessionById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoSessionsCollection)

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		return err
	}
	return nil
}
