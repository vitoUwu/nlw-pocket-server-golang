package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PendingGoal struct {
	ID                     string    `json:"id"`
	Title                  string    `json:"title"`
	DesiredWeeklyFrequency int       `json:"desiredWeeklyFrequency"`
	CreatedAt              time.Time `json:"createdAt"`
	CompletionCount        int       `json:"completionCount"`
}

func GetPendingGoals() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetString("userId")

		goalsCreatedUpToWeek, err := Db.GetUserGoalsCreatedUpToWeek(userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		goalsCompletionsCount, err := Db.CountGoalsCompletionsCreatedUpToWeek()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		pendingGoals := []PendingGoal{}
		for _, goal := range goalsCreatedUpToWeek {
			completionCount := 0
			for _, completion := range goalsCompletionsCount {
				if completion.GoalId == goal.ID {
					completionCount = completion.Count
					break
				}
			}

			pendingGoals = append(pendingGoals, PendingGoal{
				ID:                     goal.ID,
				Title:                  goal.Title,
				DesiredWeeklyFrequency: goal.DesiredWeeklyFrequency,
				CreatedAt:              goal.CreatedAt,
				CompletionCount:        completionCount,
			})
		}

		ctx.JSON(http.StatusOK, pendingGoals)
	}
}
