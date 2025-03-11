package main

import (
	"fmt"
	"strconv"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/coregate/tickets-app/configs"
	"github.com/coregate/tickets-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := configs.NewConfigs()
	server := gin.New()

	routes.RegisterRoutes(server)
	server.GET("/api/docs", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Tickets API Reference",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Fprintln(c.Writer, htmlContent)
	})

	server.Run(":" + strconv.Itoa(int(configs.App.Port)))
}
