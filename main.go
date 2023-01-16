package main

import (
	"go-gin-gorm/Player"

	"github.com/gin-gonic/gin"
)

func init() {
	Player.LoadEnvVariables()
	Player.ConnectDB()
}

func main() {

	playerRepository := Player.NewRepository(Player.DB)
	playerService := Player.NewService(playerRepository)
	playerHandler := Player.NewPlayerHandler(playerService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/players", playerHandler.GetPlayers)
	v1.GET("/player/:id", playerHandler.GetPlayerById)
	v1.POST("/player", playerHandler.CreatePlayer)
	v1.PUT("/player/:id", playerHandler.UpdatePlayer)
	v1.DELETE("/player/:id", playerHandler.DeletePlayer)

	router.Run(":8888")
}
