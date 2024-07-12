package routes

import (
	"github.com/gin-gonic/gin"
	"sample-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	// events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticated)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegisterForEvent)

	//server.POST("/events", middlewares.Authenticated, createEvent)

	// users

	server.POST("/signup", createUser)
	server.POST("/login", login)

}
