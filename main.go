package main

import (
	"github.com/coregate/tickets-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	routes.RegisterRoutes(server)

	server.Run(":8000")
}
