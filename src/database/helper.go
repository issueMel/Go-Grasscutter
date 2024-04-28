package database

import (
	"Go-Grasscutter/src/game"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAccountName(username string) *game.Account {
	account := &game.Account{}
	db.Collection("account").FindOne(context.Background(), bson.D{{"username", username}}).Decode(account)
	return account
}
