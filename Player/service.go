package Player

type Service interface {
	FindAll() ([]Players, error)
	FindById(ID int) (Players, error)
	Create(playersRequest PlayersRequest) (Players, error)
	Update(ID int, playersRequest PlayersRequest) (Players, error)
	Delete(ID int) (Players, error)
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

func (s *service) FindById(ID int) (Players, error) {
	player, err := s.repository.FindById(ID)
	return player, err
}

func (s *service) Create(playersRequest PlayersRequest) (Players, error) {
	age, _ := playersRequest.Age.Int64()
	speed, _ := playersRequest.Speed.Int64()
	power, _ := playersRequest.Power.Int64()
	players := Players{
		Name:        playersRequest.Name,
		Club:        playersRequest.Club,
		Nationality: playersRequest.Nationality,
		Boots:       playersRequest.Boots,
		Age:         int(age),
		Speed:       int(speed),
		Power:       int(power),
	}

	newPlayers, err := s.repository.Create(players)

	return newPlayers, err
}

func (s *service) Update(ID int, playersRequest PlayersRequest) (Players, error) {
	player, err := s.repository.FindById(ID)

	age, _ := playersRequest.Age.Int64()
	speed, _ := playersRequest.Speed.Int64()
	power, _ := playersRequest.Power.Int64()

	player.Name = playersRequest.Name
	player.Club = playersRequest.Club
	player.Nationality = playersRequest.Nationality
	player.Boots = playersRequest.Boots
	player.Age = int(age)
	player.Speed = int(speed)
	player.Power = int(power)

	newPlayers, err := s.repository.Update(player)

	return newPlayers, err
}

func (s *service) Delete(ID int) (Players, error) {
	player, err := s.repository.FindById(ID)
	newPlayers, err := s.repository.Delete(player)
	return newPlayers, err
}
