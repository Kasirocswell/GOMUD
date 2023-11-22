// equipment.go

package models

import (
	"bufio"
	"fmt"
	"reflect"
)

// Equipment struct defining different equipment slots
type Equipment struct {
	Right_Hand *Item
	Left_Hand  *Item
	Head       *Item
	Neck       *Item
	Chest      *Item
	Back       *Item
	Arms       *Item
	Waist      *Item
	Legs       *Item
	Feet       *Item
	Accessory  *Item
}

// equipment.go

// equipment.go

func (u *User) EquipItem(itemName string) error {
	var itemIndex int = -1

	// Find the item in the inventory by name
	for i, item := range u.Inventory {
		if item.Name == itemName {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		return fmt.Errorf("item '%s' not found in inventory", itemName)
	}

	// Determine the equipment slot for the item
	var slot **Item
	switch u.Inventory[itemIndex].Slot {
	case "Right_Hand":
		slot = &u.Equipment.Right_Hand
	case "Left_Hand":
		slot = &u.Equipment.Left_Hand
	case "Head":
		slot = &u.Equipment.Head
	case "Neck":
		slot = &u.Equipment.Neck
	case "Chest":
		slot = &u.Equipment.Chest
	case "Back":
		slot = &u.Equipment.Back
	case "Arms":
		slot = &u.Equipment.Arms
	case "Waist":
		slot = &u.Equipment.Waist
	case "Legs":
		slot = &u.Equipment.Legs
	case "Feet":
		slot = &u.Equipment.Feet
	case "Accessory":
		slot = &u.Equipment.Accessory
	default:
		return fmt.Errorf("invalid equipment slot: %s", u.Inventory[itemIndex].Slot)
	}

	// Check if the slot is already occupied
	if *slot != nil {
		return fmt.Errorf("equipment slot '%s' is already occupied", u.Inventory[itemIndex].Slot)
	}

	// Equip the item and update user's current weight
	*slot = &u.Inventory[itemIndex]
	u.CurrentWeight -= u.Inventory[itemIndex].Weight

	// Remove the item from inventory
	u.Inventory = append(u.Inventory[:itemIndex], u.Inventory[itemIndex+1:]...)

	return nil
}

// equipment.go

func (u *User) UnequipItem(itemName string) error {
	var itemPtr **Item
	var found bool

	// Check each equipment slot for the item
	val := reflect.ValueOf(&u.Equipment).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if item, ok := field.Interface().(**Item); ok && *item != nil && (*item).Name == itemName {
			itemPtr = item
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("item '%s' not found in equipment", itemName)
	}

	// Check if the user can carry more weight
	if u.CurrentWeight+(*itemPtr).Weight > u.MaxCarryWeight {
		return fmt.Errorf("cannot unequip item: will exceed carrying capacity")
	}

	// Add the item back to inventory and adjust weight
	u.Inventory = append(u.Inventory, **itemPtr)
	u.CurrentWeight += (*itemPtr).Weight

	// Clear the slot
	*itemPtr = nil

	return nil
}

func (eq *Equipment) ListEquipment(writer *bufio.Writer) {
	equippedItems := []string{
		formatEquipmentItem("Right Hand", eq.Right_Hand),
		formatEquipmentItem("Left Hand", eq.Left_Hand),
		formatEquipmentItem("Head", eq.Head),
		formatEquipmentItem("Neck", eq.Neck),
		formatEquipmentItem("Chest", eq.Chest),
		formatEquipmentItem("Back", eq.Back),
		formatEquipmentItem("Arms", eq.Arms),
		formatEquipmentItem("Waist", eq.Waist),
		formatEquipmentItem("Legs", eq.Legs),
		formatEquipmentItem("Feet", eq.Feet),
		formatEquipmentItem("Accessory", eq.Accessory),
	}

	for _, itemDesc := range equippedItems {
		if itemDesc != "" {
			writer.WriteString(itemDesc + "\n")
		}
	}

	writer.Flush()
}

func formatEquipmentItem(slotName string, item *Item) string {
	if item != nil {
		return fmt.Sprintf("%s: %s", slotName, item.Name)
	}
	return ""
}

// Additional equipment-related functions can be added here.
