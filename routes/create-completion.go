package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateGoalCompletionBody struct {
	GoalId string `json:"goalId"`
}

func CreateCompletion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetString("userId")
		var body CreateGoalCompletionBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		goalCompletionCount, err := Db.CountGoalUpToWeekCompletions(body.GoalId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var goal Goal
		err = Db.Gorm.Model(&Goal{}).
			Where("id = ? AND user_id = ?", body.GoalId, userId).
			First(&goal).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
			return
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if goal.DesiredWeeklyFrequency <= goalCompletionCount {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Goal already completed"})
			return
		}

		completion, err := Db.CreateGoalCompletion(body.GoalId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, completion)
	}
}
