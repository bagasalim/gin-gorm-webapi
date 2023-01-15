package Player

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Players struct {
	gorm.Model
	Name        string `json:"Nama" binding:"required"`
	Nationality string
	Club		string
	Age         json.Number `json:"age" binding:"required, number"`
	Speed		json.Number
	Power		json.Number
	Boots		string
}