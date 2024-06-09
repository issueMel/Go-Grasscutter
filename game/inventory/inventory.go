package inventory

import "Go-Grasscutter/data"

type Inventory struct {
	Uid            int
	Store          map[int64]*GameItem
	InventoryTypes map[int]*InventoryTab
	loaded         bool
}

func (inventory *Inventory) LoadFromDatabase() {
	if inventory.loaded {
		return
	}

	// todo Wait for avatars to load.

	items := GetInventoryItems(inventory.Uid)
	if items == nil {
		return
	}

	for _, item := range items {
		if item.Id.IsZero() {
			continue
		}

		itemData, ok := data.GameData.ItemDataMap[item.ItemId]
		if !ok {
			continue
		}

		item.ItemData = itemData

		if item.ItemData != nil {
			// todo INCOMPLETE
		}

		inventory.Store[item.Guid] = item

		if item.EquipCharacter > 0 {
			// todo INCOMPLETE
		}

		// todo INCOMPLETE
		// Load avatars after inventory.
		inventory.loaded = true
	}
}
