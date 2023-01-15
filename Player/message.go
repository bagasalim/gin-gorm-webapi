package Player

import (
	"encoding/json"
)

type PlayersRequest struct {
	Name        string `json:"Nama" binding:"required"`
	Nationality string `json:"Nationality" binding:"required"`
	Club        string `json:"Club" binding:"required"`
	Age         json.Number `json:"Age" binding:"required, number"`
	Speed       json.Number `json:"Speed" binding:"required"`
	Power       json.Number `json:"Power" binding:"required"`
	Boots       string `json:"Boots" binding:"required"`
}