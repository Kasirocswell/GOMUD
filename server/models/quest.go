// quest.go

package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Quest defines the structure of quests or missions in the game.
type Quest struct {
	ID          string
	Name        string
	Description string
	Objectives  []QuestObjective
	Reward      QuestReward
	Status      QuestStatus
}

type QuestReward struct {
	XP      int
	Items   []Item
	Credits int
}

type QuestObjectiveType string

type QuestObjective struct {
	Description string
	Type        string
	Target      string
	Current     int // Current progress (e.g., number of items collected)
	Required    int // Required progress (e.g., total items to collect)
	Completed   bool
}

type QuestStatus string

const (
	QuestNotStarted QuestStatus = "not started"
	QuestInProgress QuestStatus = "in progress"
	QuestCompleted  QuestStatus = "completed"
)

const (
	ObjectiveGather QuestObjectiveType = "gather"
	ObjectiveKill   QuestObjectiveType = "kill"
	// ... other types ...
)

var AllQuests = make(map[string]Quest)

func AssignQuestToPlayer(user *User, questID string) error {
	// Assume you have a global quest map or similar storage
	quest, exists := AllQuests[questID]
	if !exists {
		return fmt.Errorf("quest not found")
	}

	user.Quests = append(user.Quests, quest)
	user.Writer.WriteString(fmt.Sprintf("Quest '%s' has been added to your journal.\n", quest.Name))
	user.Writer.Flush()
	return nil
}

func LoadQuests(folderPath string) error {
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

		var quests []Quest
		if err := json.Unmarshal(fileData, &quests); err != nil {
			return err
		}

		for _, quest := range quests {
			AllQuests[quest.ID] = quest
		}
	}

	return nil
}

func (user *User) ListQuests() {
	if len(user.Quests) == 0 {
		user.Writer.WriteString("You have no active quests.\n")
	} else {
		user.Writer.WriteString("Your Active Quests:\n")
		for _, quest := range user.Quests {
			user.Writer.WriteString(fmt.Sprintf(" - %s: %s\n", quest.Name, quest.Status))
		}
	}
	user.Writer.Flush()
}

func (user *User) UpdateQuestObjective(questID, objectiveType, target string, amount int) {
	for i, quest := range user.Quests {
		if quest.ID == questID {
			for j, obj := range quest.Objectives {
				if obj.Type == objectiveType && obj.Target == target && !obj.Completed {
					user.Quests[i].Objectives[j].Current += amount
					if user.Quests[i].Objectives[j].Current >= obj.Required {
						user.Quests[i].Objectives[j].Completed = true
						// Check if all objectives are completed
						if user.CheckQuestCompletion(&user.Quests[i]) {
							CompleteQuest(user, quest.ID)
						}
					}
					return
				}
			}
		}
	}
}

func (user *User) CheckQuestCompletion(quest *Quest) bool {
	for _, obj := range quest.Objectives {
		if !obj.Completed {
			return false
		}
	}
	return true
}

func CompleteQuest(user *User, questID string) {
	for i, quest := range user.Quests {
		if quest.ID == questID {
			user.Quests[i].Status = QuestCompleted
			// Grant rewards
			user.XP += quest.Reward.XP
			user.Credits += quest.Reward.Credits
			// Add items to inventory, etc.
			user.Writer.WriteString(fmt.Sprintf("Quest '%s' completed! Rewards: %d XP, %d Credits\n", quest.Name, quest.Reward.XP, quest.Reward.Credits))
			user.Writer.Flush()
		}
	}
}
