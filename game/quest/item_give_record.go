package quest

type ItemGiveRecord struct {
	GivingId    int         `bson:"givingId"`
	ConfigId    int         `bson:"configId"`
	GroupId     int         `bson:"groupId"`
	LastGroupId int         `bson:"lastGroupId"`
	Finished    bool        `bson:"finished"`
	GivenItems  map[int]int `bson:"givenItems"`
}
