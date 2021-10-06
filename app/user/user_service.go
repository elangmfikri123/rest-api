package user

type Services interface {
	CreateUser(req RequestUser) (User, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreateUser(req RequestUser) (User, error) {
	user := User{}
	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
