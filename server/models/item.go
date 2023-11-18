// item.go

package models

import (
	"fmt"
	"reflect"
)

// Item represents objects that players can find, use, or trade.
type Item struct {
	ID          string
	Name        string
	Description string
	Type        string
	Stats       map[string]int
	Quantity    int
	Weight      int
	Slot        string
	Atk         int
	Def         int
	Effect      Effects
	Buff        int
}

// Add items to the inventory
func (u *User) AddItemToInventory(item Item) error {
	totalWeight := u.CurrentWeight + (item.Weight * item.Quantity)
	if totalWeight > u.MaxCarryWeight {
		return fmt.Errorf("cannot add item: carrying too much weight")
	}
	u.Inventory = append(u.Inventory, item)
	u.CurrentWeight = totalWeight
	return nil
}

// Remove items form inventory
func (u *User) RemoveItemFromInventory(itemID string) error {
	for i, item := range u.Inventory {
		if item.ID == itemID {
			u.CurrentWeight -= item.Weight * item.Quantity
			u.Inventory = append(u.Inventory[:i], u.Inventory[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found in inventory")
}

// List the user inventory
func (u *User) ListInventory() []string {
	var itemList []string

	// List items in inventory
	itemList = append(itemList, "Inventory:")
	if len(u.Inventory) == 0 {
		itemList = append(itemList, "  No items in inventory.")
	} else {
		for _, item := range u.Inventory {
			itemList = append(itemList, fmt.Sprintf("  %s: %d units, Weight: %d", item.Name, item.Quantity, item.Weight))
		}
	}

	// List equipped items
	itemList = append(itemList, "\nEquipped Items:")
	equippedItems := u.ListEquippedItems()
	if len(equippedItems) == 0 {
		itemList = append(itemList, "  No items equipped.")
	} else {
		itemList = append(itemList, equippedItems...)
	}

	return itemList
}

func (u *User) ListEquippedItems() []string {
	var equippedList []string
	val := reflect.ValueOf(u.Equipment)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if item, ok := field.Interface().(*Item); ok && item != nil {
			equippedList = append(equippedList, fmt.Sprintf("  %s: %s", val.Type().Field(i).Name, item.Name))
		}
	}
	return equippedList
}

// NewItem creates a new item with specified attributes.
func NewItem(id, name, description, itemType, slot string, stats map[string]int, quantity, weight int) *Item {
	return &Item{
		ID:          id,
		Name:        name,
		Description: description,
		Type:        itemType,
		Slot:        slot,
		Stats:       stats,
		Quantity:    quantity,
		Weight:      weight,
	}
}
