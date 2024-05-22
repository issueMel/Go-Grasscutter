package player

type PlayerCollectionRecords struct {
	Records map[int]CollectionRecord `bson:"records"`
}

type CollectionRecord struct {
	ConfigId    int   `bson:"configId"`
	ExpiredTime int64 `bson:"expiredTime"`
}
