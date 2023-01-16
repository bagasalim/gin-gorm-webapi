package Player

import "encoding/json"

type PlayersRequest struct {
	Name        string      `json:"Name" binding:"required"`
	Club        string      `json:"Club" binding:"required"`
	Nationality string      `json:"Nationality" binding:"required"`
	Age         json.Number `json:"Age" binding:"required, number"`
	Speed       json.Number `json:"Speed" binding:"required, number"`
	Power       json.Number `json:"Power" binding:"required, number"`
	Boots       string      `json:"Boots" binding:"required"`
}
