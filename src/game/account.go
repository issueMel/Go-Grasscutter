package game

import (
	"Go-Grasscutter/src/db"
	"Go-Grasscutter/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Account struct {
	ID             string   `bson:"_id"`
	Username       string   `bson:"username" index:"unique"`
	ReservedPlayer int      `bson:"reservedPlayerId"`
	Email          string   `bson:"email"`
	Token          string   `bson:"token"`
	SessionKey     string   `bson:"sessionKey"`
	Permissions    []string `bson:"permissions"`
	Locale         string   `bson:"locale"`
	BanReason      string   `bson:"banReason"`
	BanEndTime     int64    `bson:"banEndTime"`
	BanStartTime   int64    `bson:"banStartTime"`
	IsBanned       bool     `bson:"isBanned"`
}

func (a *Account) GetEmail() string {
	if len(a.Email) != 0 {
		return a.Email
	}
	// As of game version 3.5+, only the email is displayed to a user.
	return a.Username + "@go-grasscutter.io"
}

func (a *Account) GenerateSessionKey() string {
	sessionKey := utils.BytesToHex(utils.CreateSessionKey(32))
	// save in db todo
	return sessionKey
}

func (a *Account) GetAccount() *Account {
	if a == nil {
		return GetAccountById(a.ID)
	}
	return a
}

func GetAccountName(username string) *Account {
	account := &Account{}
	err := db.DB.Collection("accounts").FindOne(context.Background(), bson.D{{"username", username}}).Decode(account)
	if err != nil {
		log.Println(err)
		return account
	}
	return account
}

func GetAccountById(uid string) *Account {
	account := &Account{}
	db.DB.Collection("accounts").FindOne(context.Background(), bson.D{{"_id", uid}}).Decode(account)
	return account
}
