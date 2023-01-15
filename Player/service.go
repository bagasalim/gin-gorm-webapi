package Player

type Service interface {
	FindAll() ([]Players, error)
	FindById(Id int) (Players, error)
	Create(players PlayersRequest) (Players, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Players, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(Id int) (Players, error) {
	player, err := s.repository.FindById(Id)
	return player, err
}