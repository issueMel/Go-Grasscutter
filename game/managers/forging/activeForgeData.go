package forging

type ActiveForgeData struct {
	ForgeId             int  `bson:"forgeId"`
	AvatarId            int  `bson:"avatarId"`
	Count               int  `bson:"count"`
	StartTime           int  `bson:"startTime"`
	ForgeTime           int  `bson:"forgeTime"`
	LastUnfinishedCount int  `bson:"lastUnfinishedCount"`
	Changed             bool `bson:"changed"`
}
