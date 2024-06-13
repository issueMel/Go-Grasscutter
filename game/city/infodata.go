package city

import "Go-Grasscutter/generated/pb"

type CityInfoData struct {
	CityId     int `bson:"cityId"`
	Level      int `bson:"level"`
	NumCrystal int `bson:"numCrystal"`
}

func (c *CityInfoData) ToProto() *pb.CityInfo {
	return &pb.CityInfo{
		Level:      uint32(c.Level),
		CrystalNum: uint32(c.NumCrystal),
		CityId:     uint32(c.CityId),
	}
}
