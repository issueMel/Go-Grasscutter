package city

type CityInfoData struct {
	CityId     int `bson:"cityId"`
	Level      int `bson:"level"`
	NumCrystal int `bson:"numCrystal"`
}
