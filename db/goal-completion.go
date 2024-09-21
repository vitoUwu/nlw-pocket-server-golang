package db

import (
	"fmt"
	"time"
)

type GoalCompletion struct {
	ID        string    `json:"id"`
	GoalId    string    `json:"goalId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GoalCompletionCount struct {
	GoalId string `json:"goalId"`
	Count  int    `json:"count"`
}

func (db *Database) CountGoalsCompletionsCreatedUpToWeek() ([]GoalCompletionCount, error) {
	var goalCompletionCount []GoalCompletionCount

	err := Db.Gorm.
		Model(&GoalCompletion{}).
		Select("goal_id, COUNT(id) AS count").
		Where("created_at >= ? AND created_at <= ?", Db.FirstDayOfWeek(), Db.LastDayOfWeek()).
		Group("goal_id").
		Scan(&goalCompletionCount).Error
	if err != nil {
		return nil, err
	}

	return goalCompletionCount, nil
}

func (db *Database) CountGoalUpToWeekCompletions(goalId string) (int, error) {
	var goalCompletionCount int64

	Db.Gorm.Model(&GoalCompletion{}).
		Select("COUNT(id) AS completionCount").
		Where("created_at >= ? AND created_at <= ? AND goal_id = ?", Db.FirstDayOfWeek(), Db.LastDayOfWeek(), goalId).
		Scan(&goalCompletionCount)

	return int(goalCompletionCount), nil
}

func (db *Database) GetGoalCompletionById(id string) (*GoalCompletion, error) {
	var goalCompletion GoalCompletion
	result := db.Gorm.First(&goalCompletion).Where("id = ?", id).Take(&goalCompletion)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 1 {
		return &goalCompletion, nil
	}

	return nil, fmt.Errorf("goalCompletion with id %s not found", id)
}

func (db *Database) CreateGoalCompletion(goalId string) (*GoalCompletion, error) {
	completionWithId := GoalCompletion{
		GoalId: goalId,
		ID:     db.GenerateId(),
	}

	result := db.Gorm.Create(&completionWithId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &completionWithId, nil
}

func (db *Database) DeleteGoalCompletion(id string) (*GoalCompletion, error) {
	completionWithId := GoalCompletion{
		ID: id,
	}

	result := db.Gorm.Model(&completionWithId).Delete(&completionWithId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &completionWithId, nil
}
