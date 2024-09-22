package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Endpoint to get week summary
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} db.WeekSummary
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /summary [get]
func GetWeekSummary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userId = ctx.GetString("userId")

		summary, err := Db.GetUserWeekSummary(userId)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, summary)
	}
}
