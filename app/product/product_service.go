package product

type Services interface {
	FindProduks(userID int) (Product, error)
	CreateProducts(req RequestProduct) error
}

type services struct {
	repository Repository
}

func NewService(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreateUser(req RequestProduct) (Product, error) {
	product := Product{}
	product.ID = req.ID
	product.NmProduk = req.NmProduk
	product.Harga = req.Harga
	product.Berat = req.Berat
	product.Deskripsi = req.Deskripsi
	product.Merek = req.Merek
	product.Jenisproduk = req.Jenisproduk
	product.Tags = req.Tags
	product.Color = req.Color
	product.Size = req.Size
	product.Bahan = req.Bahan
	product.Stok = req.Stok
	product.Gambar = req.Gambar
	product.UserID = req.UserID

	newProduct, err := s.repository.InsertProduct(product)
	

	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

// func (s *service) FindProduks(userID int) (Product, error) {
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
