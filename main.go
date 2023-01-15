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
	router := gin.Default()

	// Test Create
	// player := Player.Players{}
	// player.Name = "Craig Bellamy"
	// player.Nationality = "France"
	// player.Club = "Manchester City"
	// player.Speed = "99"
	// player.Power = "99"
	// player.Boots = "Mizuno Accelerator"

	// err := Player.DB.Create(&player).Error
	// if err != nil {
	// 	fmt.Println("Err: ")
	// }

	// Test Get
	// var player []Player.Players
	// err := Player.DB.Find(&player).Error
	// if err != nil {
	// 		fmt.Println("Err: ")
	// }

	// for _, p := range player {
	// 	fmt.Println("Nama: ", p.Name)
	// fmt.Println("Boots: ",p.Boots)
	// }

	// Test Update
	// var player Player.Players
	// err := Player.DB.Debug().Where("id = ?", 1).First(&player).Error
	// if err != nil {
	// 	fmt.Println("Err: ")
	// }

	// player.Name = "Badiashile"
	// err = Player.DB.Save(&player).Error
	// if err != nil {
	// 	fmt.Println("Err: ")
	// }

	// Test Delete
	// var player Player.Players
	// err := Player.DB.Debug().Where("id = ?", 1).First(&player).Error
	// if err != nil {
	// 	fmt.Println("Err: ")
	// }

	// err = Player.DB.Delete(&player).Error
	// if err != nil {
	// 	fmt.Println("Err: ")
	// }

	playerRepository := Player.NewRepository(Player.DB)

	// Players, _ := playerRepository.FindAll()
	
	// for _, player := range Players {
	// 	fmt.Println("Boots: ", player.Boots)
	// }

	// player, _ := playerRepository.FindById(2) 
	// fmt.Println("Name: ", player.Name)
	
	Players := Player.Players{
		Name: "Red",
		Nationality: "Meng",
		Club: "Tom",
		Speed: "99",
		Power: "99",
		Boots: "Nike Phantom",
	}

	playerRepository.Create(Players)

	v1 := router.Group("/v1")

	v1.GET("/", Player.RootHandler)	
	v1.GET("/hello", Player.HelloHandler)	
	v1.GET("/player/:id/:title", Player.IdPlayerHandler)	
	v1.GET("/query", Player.QueryHandler)	
	v1.POST("/player", Player.CreatePlayer)
	
	
	router.Run(":8888")
}