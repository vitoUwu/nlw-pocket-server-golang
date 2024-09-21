package middlewares

import (
	"fmt"
	"net/http"
	"nlw/pocket/db"

	"github.com/gin-gonic/gin"
)

var (
	Db = db.Db
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := ctx.Cookie("userId")

		if err != nil || userId == "" {
			user, err := Db.CreateUser()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}
			userId = user.ID
			ctx.Header("Set-Cookie", fmt.Sprintf("userId=%s; Max-Age=2592000; Path=/; Secure; HttpOnly; SameSite=none", userId))
			ctx.Set("userId", userId)

			ctx.Next()

			return
		}

		_, err = Db.GetUserById(userId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "userId is invalid. Clear your cookies and try again."})
			ctx.Abort()
			return
		}

		ctx.Set("userId", userId)
		ctx.Next()
	}
}
