package inventory

type Inventory struct {
	Uid            int
	Store          map[int64]*GameItem
	InventoryTypes map[int]*InventoryTab
}

func LoadFromDatabase() {

}
