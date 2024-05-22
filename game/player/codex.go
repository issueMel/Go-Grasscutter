package player

type PlayerCodex struct {
	UnlockedWeapon             []int       `bson:"unlockedWeapon"`
	UnlockedAnimal             map[int]int `bson:"unlockedAnimal"`
	unlockedMaterial           []int       `bson:"unlockedMaterial"`
	unlockedBook               []int       `bson:"unlockedBook"`
	unlockedTip                []int       `bson:"unlockedTip"`
	unlockedView               []int       `bson:"unlockedView"`
	unlockedReliquary          []int       `bson:"unlockedReliquary"`
	unlockedReliquarySuitCodex []int       `bson:"unlockedReliquarySuitCodex"`
}
