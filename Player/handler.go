package Player

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type playerHandler struct {
	playerService Service
}

func NewPlayerHandler(playerService Service) *playerHandler {
	return &playerHandler{playerService}
}

func (h *playerHandler) GetPlayers(c *gin.Context) {
	players, err := h.playerService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var playersResponse []PlayersResponse
	for _, p := range players {

		playerResponse := convertToPlayerResponse(p)

		playersResponse = append(playersResponse, playerResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": playersResponse,
	})
}

func (h *playerHandler) GetPlayerById(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	p, err := h.playerService.FindById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	playersResponse := convertToPlayerResponse(p)
	c.JSON(http.StatusOK, gin.H{
		"data": playersResponse,
	})

}

func (h *playerHandler) CreatePlayer(c *gin.Context) {
	var playersRequest PlayersRequest

	err := c.ShouldBindJSON(&playersRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return

	}

	player, err := h.playerService.Create(playersRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": player,
	})

}

func (h *playerHandler) UpdatePlayer(c *gin.Context) {
	var playersRequest PlayersRequest

	err := c.ShouldBindJSON(&playersRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return

	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	player, err := h.playerService.Update(id, playersRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToPlayerResponse(player),
	})

}

func (h *playerHandler) DeletePlayer(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	p, err := h.playerService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	playersResponse := convertToPlayerResponse(p)
	c.JSON(http.StatusOK, gin.H{
		"data": playersResponse,
	})

}

func convertToPlayerResponse(p Players) PlayersResponse {
	return PlayersResponse{
		ID:          p.ID,
		Name:        p.Name,
		Club:        p.Club,
		Nationality: p.Nationality,
		Boots:       p.Boots,
		Power:       p.Power,
		Speed:       p.Speed,
		Age:         p.Age,
	}
}

// Test Create
// player = Players{}
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
// err = Player.DB.Find(&player).Error
// if err != nil {
// 		fmt.Println("Err: ")
// }

// for _, p := range player {
// 	fmt.Println("Nama: ", p.Name)
// fmt.Println("Boots: ",p.Boots)
// }

// Test Update
// var player Player.Players
// err = Player.DB.Debug().Where("id = ?", 1).First(&player).Error
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
// err = Player.DB.Debug().Where("id = ?", 1).First(&player).Error
// if err != nil {
// 	fmt.Println("Err: ")
// }

// err = Player.DB.Delete(&player).Error
// if err != nil {
// 	fmt.Println("Err: ")
// }
