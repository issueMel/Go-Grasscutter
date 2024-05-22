package gacha

type PlayerGachaBannerInfo struct {
	TotalPulls               int `bson:"totalPulls"`
	Pity5                    int `bson:"pity5"`
	Pity4                    int `bson:"pity4"`
	FailedFeaturedItemPulls  int `bson:"failedFeaturedItemPulls"`
	FailedFeatured4ItemPulls int `bson:"failedFeatured4ItemPulls"`
	Pity5Pool1               int `bson:"pity5Pool1"`
	Pity5Pool2               int `bson:"pity5Pool2"`
	Pity4Pool1               int `bson:"pity4Pool1"`
	Pity4Pool2               int `bson:"pity4Pool2"`
	FailedChosenItemPulls    int `bson:"failedChosenItemPulls"`
	WishItemId               int `bson:"wishItemId"`
}
