package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) 
	server.POST("/events", createEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
