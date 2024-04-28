package game

import "Go-Grasscutter/utils"

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
	if len(a.Email) == 0 {
		return a.Email
	}
	// As of game version 3.5+, only the email is displayed to a user.
	return a.Username + "@go-grasscutter.io"
}

func (a *Account) GenerateSessionKey() string {
	sessionKey := utils.BytesToHex(utils.CreateSessionKey(32))
	// save in database todo
	return sessionKey
}
