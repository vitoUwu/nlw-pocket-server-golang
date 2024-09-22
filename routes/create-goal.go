package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	db "nlw/pocket/db"
)

type CreateGoalBody struct {
	Title                  string `json:"title" binding:"required"`
	DesiredWeeklyFrequency int    `json:"desiredWeeklyFrequency" binding:"required"`
}

type CreateGoalResponse struct {
	Title                  string `json:"title"`
	DesiredWeeklyFrequency int    `json:"desiredWeeklyFrequency"`
}

// @Summary Endpoint to create a goal
// @Schemes
// @Accept json
// @Produce json
// @Success 201 {object} CreateGoalResponse
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Param body body CreateGoalBody true "body params"
// @Router /goals [post]
func CreateGoal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := CreateGoalBody{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId := ctx.GetString("userId")
		goal, err := Db.CreateGoal(db.CreateGoal{
			Title:                  body.Title,
			DesiredWeeklyFrequency: body.DesiredWeeklyFrequency,
			UserId:                 userId,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"title":                  goal.Title,
			"desiredWeeklyFrequency": goal.DesiredWeeklyFrequency,
		})
	}
}
