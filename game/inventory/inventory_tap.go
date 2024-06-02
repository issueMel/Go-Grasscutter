package inventory

type InventoryTab interface {
	GetItemById(id int) *GameItem
	OnAddItem(item *GameItem)
	OnRemoveItem(item *GameItem)
	GetSize() int
	GetMaxCapacity() int
	GetItemCountById(itemId int) int
}
