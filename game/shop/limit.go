package shop

type ShopLimit struct {
	ShopGoodId        int `bson:"shopGoodId"`
	HasBought         int `bson:"hasBought"`
	HasBoughtInPeriod int `bson:"hasBoughtInPeriod"`
	NextRefreshTime   int `bson:"nextRefreshTime"`
}
