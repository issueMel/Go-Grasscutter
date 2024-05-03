package game

import (
	"Go-Grasscutter/src/db"
	utils2 "Go-Grasscutter/src/utils"
	"Go-Grasscutter/src/utils/crypto"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (a *Account) SaveAccount() {
	// todo async store data
	db.DB.Collection("accounts").ReplaceOne(context.Background(), bson.M{"_id": a.ID}, a)
}

func (a *Account) GetEmail() string {
	if len(a.Email) != 0 {
		return a.Email
	}
	// As of game version 3.5+, only the email is displayed to a user.
	return a.Username + "@go-grasscutter.io"
}

func (a *Account) GenerateSessionKey() string {
	sessionKey := utils2.BytesToHex(crypto.CreateSessionKey(32))
	// save in db
	a.SessionKey = sessionKey
	a.SaveAccount()
	return sessionKey
}

func (a *Account) GenerateLoginToken() string {
	token := utils2.BytesToHex(crypto.CreateSessionKey(32))
	// save in db
	a.Token = token
	a.SaveAccount()
	return token
}

func (a *Account) GetAccount() *Account {
	if a == nil {
		return GetAccountById(a.ID)
	}
	return a
}

func GetAccountName(username string) *Account {
	account := &Account{}
	db.DB.Collection("accounts").FindOne(context.Background(), bson.D{{"username", username}}).Decode(account)
	return account
}

func GetAccountById(uid string) *Account {
	account := &Account{}
	db.DB.Collection("accounts").FindOne(context.Background(), bson.D{{"_id", uid}}).Decode(account)
	return account
}
