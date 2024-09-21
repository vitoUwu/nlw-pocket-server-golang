package main

import (
	middlewares "nlw/pocket/middewares"
	routes "nlw/pocket/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	authenticated := router.Group("/")

	authenticated.Use(middlewares.AuthMiddleware())
	{
		authenticated.POST("/goals", routes.CreateGoal())
		authenticated.GET("/pending-goals", routes.GetPendingGoals())
		authenticated.GET("/summary", routes.GetWeekSummary())
		authenticated.POST("/completions", routes.CreateCompletion())
		authenticated.DELETE("/completions", routes.DeleteCompletion())
	}

	router.Run(":8000")
}
