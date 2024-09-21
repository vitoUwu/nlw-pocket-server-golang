package db

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Goal struct {
	ID                     string            `json:"id"`
	Title                  string            `json:"title"`
	DesiredWeeklyFrequency int               `json:"desiredWeeklyFrequency"`
	UserId                 string            `json:"userId"`
	Completions            []*GoalCompletion `json:"completions" gorm:"foreignKey:GoalId;references:ID;OnDelete:CASCADE"`
	CreatedAt              time.Time         `json:"createdAt"`
	UpdatedAt              time.Time         `json:"updatedAt"`
}

type CreateGoal struct {
	Title                  string `json:"title"`
	DesiredWeeklyFrequency int    `json:"desiredWeeklyFrequency"`
	UserId                 string `json:"userId"`
}

func (db *Database) CreateGoal(goal CreateGoal) (*Goal, error) {
	goalWithId := Goal{
		ID:                     db.GenerateId(),
		Title:                  goal.Title,
		UserId:                 goal.UserId,
		DesiredWeeklyFrequency: goal.DesiredWeeklyFrequency,
	}

	result := db.Gorm.Create(&goalWithId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &goalWithId, nil
}

func (db *Database) GetGoalById(id string, userId string) (*Goal, error) {
	var goal Goal
	result := db.Gorm.First(&goal).Where("id = ? AND user_id = ?", id, userId).Take(&goal)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 1 {
		return &goal, nil
	}

	return nil, fmt.Errorf("goal with id %s not found", id)
}

func (db *Database) GetGoalsByUserId(userId string) ([]Goal, error) {
	goals := []Goal{}
	result := db.Gorm.Model(&Goal{}).Where("user_id = ?", userId).Find(&goals)
	if result.Error != nil {
		return nil, result.Error
	}
	return goals, nil
}

func (db *Database) GetUserGoalsCreatedUpToWeek(userId string) ([]Goal, error) {
	goals := []Goal{}
	err := db.Gorm.Model(&Goal{}).Where("created_at <= ? AND user_id = ?", Db.LastDayOfWeek(), userId).Find(&goals).Error
	if err != nil {
		return nil, err
	}
	return goals, nil
}

type SummaryGoal struct {
	ID          string `json:"id"`
	GoalID      string `json:"goalId"`
	Title       string `json:"title"`
	CompletedAt string `json:"completedAt"`
}

type WeekSummary struct {
	Completed   int64                    `json:"completed"`
	Total       int64                    `json:"total"`
	GoalsPerDay map[string][]SummaryGoal `json:"goalsPerDay"`
}

func (db *Database) GetUserWeekSummary(userId string) (*WeekSummary, error) {
	var rawSummary map[string]interface{}

	firstDayOfWeek := db.FirstDayOfWeek()
	lastDayOfWeek := db.LastDayOfWeek()

	Db.Gorm.Raw(`
			SELECT 
				(
					SELECT COUNT(*)
					FROM (
						SELECT goal_completions.id
						FROM goal_completions
						INNER JOIN goals ON goal_completions.goal_id = goals.id
						WHERE goal_completions.created_at BETWEEN ? AND ?
						AND goals.user_id = ?
					) AS goals_completed_in_week
				) AS completed,

				(
					SELECT CAST(SUM(desired_weekly_frequency) AS BIGINT)
					FROM goals
					WHERE created_at <= ?
					AND user_id = ?
				) AS total,

				(
					SELECT JSON_OBJECT_AGG(completed_at_date, completions)
					FROM (
						SELECT
							DATE(goal_completions.created_at) AS completed_at_date,
							JSON_AGG(
								JSON_BUILD_OBJECT(
									'id', goal_completions.id,
									'goalId', goal_completions.goal_id,
									'title', goals.title,
									'completedAt', goal_completions.created_at
								)
							) AS completions
						FROM goal_completions
						INNER JOIN goals ON goal_completions.goal_id = goals.id
						WHERE goal_completions.created_at BETWEEN ? AND ?
						AND goals.user_id = ?
						GROUP BY completed_at_date
						ORDER BY completed_at_date DESC
					) AS goals_completed_by_week_day
				) AS goalsPerDay

			`, firstDayOfWeek, lastDayOfWeek, userId, lastDayOfWeek, userId, firstDayOfWeek, lastDayOfWeek, userId).
		Scan(&rawSummary)

	var goalsPerDay map[string][]SummaryGoal

	if val, ok := rawSummary["goalsperday"]; ok && val != nil {
		goalsPerDayStr := val.(string)

		err := json.Unmarshal([]byte(goalsPerDayStr), &goalsPerDay)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		goalsPerDay = map[string][]SummaryGoal{}
	}

	completed, ok := rawSummary["completed"].(int64)
	if !ok {
		completed = 0
	}

	total, ok := rawSummary["total"].(int64)
	if !ok {
		total = 0
	}

	return &WeekSummary{
		Completed:   completed,
		Total:       total,
		GoalsPerDay: goalsPerDay,
	}, nil
}
