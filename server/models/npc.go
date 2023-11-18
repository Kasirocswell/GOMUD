package models

type NPC struct {
	ID          string // Unique identifier for the NPC
	Name        string // Name of the NPC
	Description string // Description of the NPC
	Race        string // Race of the NPC, if applicable
	Class       string // Class or profession of the NPC

	Health     int            // Health points, for NPCs involved in combat
	MaxHealth  int            // Maximum health points
	Attributes Attributes     // NPC's attributes (strength, intelligence, etc.)
	Skills     map[string]int // Skills the NPC possesses

	Inventory []Item    // Items that the NPC carries
	Equipment Equipment // Equipment the NPC is wearing

	AIType   AIType   // Type of AI (hostile, friendly, merchant, etc.)
	Dialogue []string // Dialogue lines for interaction
	Quests   []string // IDs of quests this NPC can give

	LocationID string `json:"LocationID"`
	Location *Room // Current location of the NPC
	// Additional fields related to NPC behavior, story role, etc.
}

// AIType defines the behavior pattern of an NPC
type AIType string

const (
	Friendly   AIType = "Friendly"
	Hostile    AIType = "Hostile"
	Shop       AIType = "Shop"
	QuestGiver AIType = "QuestGiver"
	// More AI types as needed
)

func InteractWithNPC(user *User, npcName string) {
	// Interaction logic...
}

// NewNPC creates a new NPC with specified attributes.
func NewNPC(name, description, race, class string, health, maxHealth int, dialogue []string, quests []string) *NPC {
	return &NPC{
		Name:        name,
		Description: description,
		Race:        race,
		Class:       class,
		Health:      health,
		MaxHealth:   maxHealth,
		Dialogue:    dialogue,
		Quests:      quests,
		// Initialize other fields as needed
	}
}
