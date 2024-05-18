package object

type ComboTokenReqJson struct {
	AppID          int    `json:"app_id"`
	ChannelID      int    `json:"channel_id"`
	Data           string `json:"data"`
	Device         string `json:"device"`
	Sign           string `json:"sign"`
	LoginTokenData LoginTokenData
}

type LoginTokenData struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
	Guest bool   `json:"guest"`
}

type ComboTokenResJson struct {
	Message string    `json:"message"`
	Retcode int       `json:"retcode"`
	Data    LoginData `json:"data"`
}

type LoginData struct {
	AccountType   int    `json:"account_type"`
	Heartbeat     bool   `json:"heartbeat"`
	ComboID       string `json:"combo_id"`
	ComboToken    string `json:"combo_token"`
	OpenID        string `json:"open_id"`
	Data          string `json:"data"`
	FatigueRemind *bool  `json:"fatigue_remind,omitempty"`
}

func NewComboTokenResJson() *ComboTokenResJson {
	return &ComboTokenResJson{
		Data: LoginData{
			AccountType:   1,
			Heartbeat:     false,
			ComboID:       "",
			ComboToken:    "",
			OpenID:        "",
			Data:          "{\"guest\":false}",
			FatigueRemind: nil,
		},
	}
}
