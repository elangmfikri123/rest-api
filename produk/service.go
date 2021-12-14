package produk

type Service interface {
	FindProduks(userID int) (Produk, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// func (s *service) FindProduks(userID int) (Produk, error) {
// 	if userID != 0 {
// 		produks, err := s.repository.FindByUserID(userID)
// 		if err != nil {
// 			return produks, err
// 		}
// 		return produks, nil
// 	}
// 	// produks, err := s.repository.FindAll()
// 	// 	if err != nil {
// 	// 		return produks, err
// 	// 	}
// 	// 	return produks, nil
// }
