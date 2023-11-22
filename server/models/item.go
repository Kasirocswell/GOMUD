// item.go

package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
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
func (u *User) RemoveItemFromInventory(itemName string) (Item, error) {
	for i, item := range u.Inventory {
		if strings.EqualFold(item.Name, itemName) { // Case-insensitive comparison
			u.CurrentWeight -= item.Weight * item.Quantity
			removedItem := u.Inventory[i]
			u.Inventory = append(u.Inventory[:i], u.Inventory[i+1:]...)
			return removedItem, nil
		}
	}
	return Item{}, fmt.Errorf("item '%s' not found in inventory", itemName)
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

var AllWeapons = make(map[string]Item)

func LoadWeapons(folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(folderPath, file.Name())
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		var weapons []Item
		if err := json.Unmarshal(fileData, &weapons); err != nil {
			return err
		}

		for _, weapon := range weapons {
			AllWeapons[weapon.ID] = weapon
		}
	}

	return nil
}

var AllArmor = make(map[string]Item)

func LoadArmor(folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(folderPath, file.Name())
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		var armorItems []Item
		if err := json.Unmarshal(fileData, &armorItems); err != nil {
			return err
		}

		for _, armor := range armorItems {
			AllArmor[armor.ID] = armor
		}
	}

	return nil
}

var AllItems = make(map[string]Item)

func LoadItems(folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(folderPath, file.Name())
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		var items []Item
		if err := json.Unmarshal(fileData, &items); err != nil {
			return err
		}

		for _, item := range items {
			AllItems[item.ID] = item
		}
	}

	return nil
}
