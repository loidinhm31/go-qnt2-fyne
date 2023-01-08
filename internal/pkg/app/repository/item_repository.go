package repository

import (
	"context"
	"go-qn2management/internal/pkg/app/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type SessionItem struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Title     string    `bson:"title" json:"title"`
	Extension string    `bson:"extension" json:"extension"`
	SessionID string    `bson:"session_id" json:"session_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (mongo *mongoRepository) FindAllItems() ([]*SessionItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoItemsCollection)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessionItems []*SessionItem

	for cursor.Next(ctx) {
		var item SessionItem

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding item into slice:", err)
			return nil, err
		} else {
			sessionItems = append(sessionItems, &item)
		}
	}
	return sessionItems, nil
}

func (mongo *mongoRepository) FindItemsBySessionId(sessionId string) ([]*SessionItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoItemsCollection)

	cursor, err := collection.Find(ctx,
		bson.M{"session_id": sessionId},
	)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessionItems []*SessionItem

	for cursor.Next(ctx) {
		var item SessionItem

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding log into slice:", err)
			return nil, err
		} else {
			sessionItems = append(sessionItems, &item)
		}
	}
	return sessionItems, nil
}

func (mongo *mongoRepository) InsertItem(sessionItem *SessionItem) error {
	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoItemsCollection)

	_, err := collection.InsertOne(context.TODO(), SessionItem{
		Title:     sessionItem.Title,
		Extension: sessionItem.Extension,
		SessionID: sessionItem.SessionID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Error inserting into sessions", err)
		return err
	}
	return nil
}

func (mongo *mongoRepository) DeleteItemById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := mongo.mongoClient.Database(config.MongoDB).Collection(config.MongoItemsCollection)

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
