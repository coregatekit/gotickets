package main

import (
	"strconv"

	"github.com/coregate/tickets-app/configs"
	"github.com/coregate/tickets-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := configs.NewConfigs()
	server := gin.New()

	routes.RegisterRoutes(server)

	server.Run(":" + strconv.Itoa(int(configs.App.Port)))
}
