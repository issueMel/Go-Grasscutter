package gacha

type PlayerGachaInfo struct {
	StandardBanner       *PlayerGachaBannerInfo `bson:"standardBanner"`
	beginnerBanner       *PlayerGachaBannerInfo `bson:"beginnerBanner"`
	eventCharacterBanner *PlayerGachaBannerInfo `bson:"eventCharacterBanner"`
	eventWeaponBanner    *PlayerGachaBannerInfo `bson:"eventWeaponBanner"`
}
