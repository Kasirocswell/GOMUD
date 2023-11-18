// quest.go

package models

// Quest defines the structure of quests or missions in the game.
type Quest struct {
	ID          string
	Name        string
	Description string
	Objectives  []string
	Reward      QuestReward
	Status      QuestStatus
}

type QuestReward struct {
	XP      int
	Items   []Item
	Credits int
}

type QuestStatus string

const (
	QuestNotStarted QuestStatus = "not started"
	QuestInProgress QuestStatus = "in progress"
	QuestCompleted  QuestStatus = "completed"
)

// Additional methods for quest management can be added here.
