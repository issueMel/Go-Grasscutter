package expedition

type ExpeditionInfo struct {
	State     int `bson:"state"`
	ExpId     int `bson:"expId"`
	HourTime  int `bson:"hourTime"`
	StartTime int `bson:"startTime"`
}
