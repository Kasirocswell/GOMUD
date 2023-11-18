package models

import "time"

type Enemy struct {
	NPC // Inherit properties from the NPC struct

	AggressionLevel int     // Indicates how aggressive the enemy is
	CombatStyle     string  // Defines the enemy's combat style or tactics
	PatrolRoute     []*Room // If the enemy patrols, a list of rooms they move through
	LootTable       []Item  // Potential items dropped by the enemy on defeat
	ExperienceValue int     // The amount of XP granted to the player upon defeating the enemy
	RespawnTime     int
	DeathTime       time.Time
	IsDead          bool
	StatusEffects   Effects
	Attack          int
	Defense         int
}

var DeadEnemies []*Enemy

// NewEnemy creates a new enemy with specified attributes.
func NewEnemy(name, description, race, class string, health, maxHealth int, aggressionLevel int, combatStyle string, lootTable []Item, xpValue int, respawnTime int) *Enemy {
	return &Enemy{
		NPC: NPC{
			Name:        name,
			Description: description,
			Race:        race,
			Class:       class,
			Health:      health,
			MaxHealth:   maxHealth,
			// Initialize other NPC fields as needed
		},
		AggressionLevel: aggressionLevel,
		CombatStyle:     combatStyle,
		LootTable:       lootTable,
		ExperienceValue: xpValue,
		RespawnTime:     respawnTime,
		// Set other enemy-specific fields
	}
}
