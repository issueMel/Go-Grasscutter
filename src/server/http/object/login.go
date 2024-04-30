package object

type LoginResultJson struct {
	Message string     `json:"message"`
	RetCode int        `json:"retcode"`
	Data    VerifyData `json:"data"`
}

type LoginAccountRequestJson struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IsCrypto bool   `json:"is_crypto"`
}

type VerifyData struct {
	Account             VerifyAccountData `json:"account"`
	DeviceGrantRequired bool              `json:"device_grant_required"`
	RealNameOperation   string            `json:"realname_operation"`
	RealPersonRequired  bool              `json:"realperson_required"`
	SafeMobileRequired  bool              `json:"safe_mobile_required"`
}

type VerifyAccountData struct {
	UID               string `json:"uid"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	IsEmailVerify     string `json:"is_email_verify"`
	RealName          string `json:"realname"`
	IdentityCard      string `json:"identity_card"`
	Token             string `json:"token"`
	SafeMobile        string `json:"safe_mobile"`
	FacebookName      string `json:"facebook_name"`
	TwitterName       string `json:"twitter_name"`
	GameCenterName    string `json:"game_center_name"`
	GoogleName        string `json:"google_name"`
	AppleName         string `json:"apple_name"`
	SonyName          string `json:"sony_name"`
	TapName           string `json:"tap_name"`
	Country           string `json:"country"`
	ReactivateTicket  string `json:"reactivate_ticket"`
	AreaCode          string `json:"area_code"`
	DeviceGrantTicket string `json:"device_grant_ticket"`
}

func NewLoginResultJson() *LoginResultJson {
	return &LoginResultJson{
		RetCode: 0,
		Data: VerifyData{
			Account: VerifyAccountData{
				IsEmailVerify: "0",
				Country:       "US",
				AreaCode:      "**",
			},
			RealNameOperation: "NONE",
		},
	}
}
