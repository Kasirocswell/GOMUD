package models

import (
	"fmt"
	"strconv"
)

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

	AIType            AIType // Type of AI (hostile, friendly, merchant, etc.)
	DialogueStartNode string
	Quests            []string // IDs of quests this NPC can give

	LocationID string `json:"LocationID"`
	Location   *Room  // Current location of the NPC
	// Additional fields related to NPC behavior, story role, etc.
}

type DialogueNode struct {
	ID        string
	Text      string
	Responses []Response
}

// Response represents a possible response in a dialogue node.
type Response struct {
	Text     string
	NextNode string // ID of the next DialogueNode
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

func InteractWithNPC(user *User, npc *NPC, dialogueNodes map[string]DialogueNode) {
	if npc.DialogueStartNode == "" {
		user.Writer.WriteString("This NPC has nothing to say.\n")
		user.Writer.Flush()
		return
	}

	currentNode := dialogueNodes[npc.DialogueStartNode]
	for {
		// Display the dialogue text
		user.Writer.WriteString(currentNode.Text + "\n")

		// Offer response choices
		for i, response := range currentNode.Responses {
			user.Writer.WriteString(fmt.Sprintf("%d. %s\n", i+1, response.Text))
		}

		// Add an option to end the conversation
		endConversationChoice := len(currentNode.Responses) + 1
		user.Writer.WriteString(fmt.Sprintf("%d. End Conversation\n", endConversationChoice))
		user.Writer.Flush()

		// Capture user input using the same method as in HandleCommand
		input, err := user.ReadInput()
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Convert input to choice number
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > endConversationChoice {
			user.Writer.WriteString("Invalid choice. Try again.\n")
			user.Writer.Flush()
			continue
		}

		// Handle choice to end conversation
		if choice == endConversationChoice {
			return
		}

		nextNodeID := currentNode.Responses[choice-1].NextNode
		if nextNodeID == "" {
			// End of dialogue
			return
		}
		currentNode = dialogueNodes[nextNodeID]
	}
}

// NewNPC creates a new NPC with specified attributes.
// func NewNPC(name, description, race, class string, health, maxHealth int, dialogue []string, quests []string) *NPC {
// 	return &NPC{
// 		Name:        name,
// 		Description: description,
// 		Race:        race,
// 		Class:       class,
// 		Health:      health,
// 		MaxHealth:   maxHealth,
// 		Dialogue:    dialogue,
// 		Quests:      quests,
// 		// Initialize other fields as needed
// 	}
// }
