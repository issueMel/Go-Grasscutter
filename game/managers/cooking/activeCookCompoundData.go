package cooking

type ActiveCookCompoundData struct {
	CostTime   int `bson:"costTime"`
	CompoundId int `bson:"compoundId"`
	TotalCount int `bson:"totalCount"`
	StartTime  int `bson:"startTime"`
}
