package routes

import (
	"fmt"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/coregate/tickets-app/configs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB, configs *configs.Configs) {
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

	HealthCheckRoutes(server)
	AuthRoutes(server, db, configs)
}
