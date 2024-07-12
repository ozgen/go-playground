package main

import (
	"github.com/gin-gonic/gin"
	"sample-rest-api/db"
	"sample-rest-api/routes"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080")
}
