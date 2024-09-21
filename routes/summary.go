package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
