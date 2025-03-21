package main

import (
	"log"
	"strconv"

	"github.com/coregate/tickets-app/configs"
	"github.com/coregate/tickets-app/database"
	"github.com/coregate/tickets-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := configs.NewConfigs()
	server := gin.Default()

	dbConnection, err := database.NewPostgres(configs)
	if err != nil {
		log.Fatalln(err)
	}

	routes.RegisterRoutes(server, dbConnection.DB, configs)

	err = server.Run(":" + strconv.Itoa(int(configs.App.Port)))
	if err != nil {
		log.Fatalln(err)
	}
}
