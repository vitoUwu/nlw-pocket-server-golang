package main

import (
	middlewares "nlw/pocket/middewares"
	routes "nlw/pocket/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://nlw-pocket-web-production.up.railway.app", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
