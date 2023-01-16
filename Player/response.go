package Player

type PlayersResponse struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"nama" binding:"required"`
	Club        string `json:"club" binding:"required"`
	Nationality string `json:"nationality" binding:"required"`
	Age         int    `json:"age" binding:"required"`
	Speed       int    `json:"speed" binding:"required"`
	Power       int    `json:"power" binding:"required"`
	Boots       string `json:"boots" binding:"required"`
}
