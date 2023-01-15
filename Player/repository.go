package Player

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Players, error)
	FindById(ID int) (Players, error)
	Create(player Players) (Players, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Players, error) {
	
	var player []Players
	
	err := r.db.Find(&player).Error

	return player, err
}

func (r *repository) FindById(ID int) (Players, error) {
	
	var player Players
	
	err := r.db.Find(&player).Error

	return player, err
}

func (r *repository) Create(player Players) (Players, error) {
	
	err := r.db.Create(&player).Error

	return player, err
}