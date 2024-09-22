package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteGoalCompletionBody struct {
	GoalId       string `json:"goalId"`
	CompletionId string `json:"completionId"`
}

type DeleteGoalCompletionResponse struct {
	Message string `json:"message"`
}

// @Summary Endpoint to delete a goal completion
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} DeleteGoalCompletionResponse
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Param body body DeleteGoalCompletionBody true "body params"
// @Router /completions [delete]
func DeleteCompletion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetString("userId")

		var body DeleteGoalCompletionBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		completion, err := Db.GetGoalCompletionById(body.CompletionId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Completion not found"})
			return
		}

		_, err = Db.GetGoalById(completion.GoalId, userId)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
			return
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err = Db.DeleteGoalCompletion(body.CompletionId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Completion deleted"})
	}
}
