package db

import (
	"Go-Grasscutter/log"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Counter struct {
	ID    string `bson:"_id"`
	Count int    `bson:"count"`
}

func (c *Counter) save() {
	_, err := DB.Collection("counters").InsertOne(context.Background(), c)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
}

func NewDatabaseCounter(id string) *Counter {
	return &Counter{
		ID:    id,
		Count: 10000,
	}
}

func GetNextId(name string) int {
	c := &Counter{}
	err := DB.Collection("counters").FindOne(context.Background(), bson.D{{"_id", name}}).Decode(c)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c = NewDatabaseCounter(name)
		}
		log.SugaredLogger.Error(err)
		return 0
	}
	result := c.GetNextId()
	go c.save()
	return result
}

func (c *Counter) GetNextId() int {
	id := c.Count
	c.Count++
	return id
}
