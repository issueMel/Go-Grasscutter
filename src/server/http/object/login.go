package object

type LoginResultJson struct {
	Message string
	RetCode int
	Data    VerifyData
}

type LoginAccountRequestJson struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IsCrypto bool   `json:"is_crypto"`
}

type VerifyData struct {
	Account             VerifyAccountData
	DeviceGrantRequired bool
	RealNameOperation   string
	RealPersonRequired  bool
	SafeMobileRequired  bool
}

type VerifyAccountData struct {
	UID               string
	Name              string
	Email             string
	Mobile            string
	IsEmailVerify     string
	RealName          string
	IdentityCard      string
	Token             string
	SafeMobile        string
	FacebookName      string
	TwitterName       string
	GameCenterName    string
	GoogleName        string
	AppleName         string
	SonyName          string
	TapName           string
	Country           string
	ReactivateTicket  string
	AreaCode          string
	DeviceGrantTicket string
}
