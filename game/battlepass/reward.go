package battlepass

type Reward struct {
	Level    int  `bson:"level"`
	RewardId int  `bson:"rewardId"`
	Paid     bool `bson:"paid"`
}
