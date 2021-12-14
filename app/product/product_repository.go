package product

import (
	"github.com/jinzhu/gorm"
	//"gorm.io/gorm"
)

type Repository interface {
	//FindAll() ([]Produk, error)
	InsertProduct(product Product) (Product, error)
	FindAll(id int) *Product
	FindByUserID(id int) (Product, error)
}

func (r *repository) InsertProduct(product Product) (Product, error) {
	// users = append(users, user)

	// fmt.Println(users)
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(id string) *Product {
	var products Product
	err := r.db.Preload("ProductImages", "product_images.is_primary = 1").Find(&products).Error
	if err == nil {
		return &products
	}

	return nil
}

func (r *repository) FindByUserID(userID int) (Product, error) {
	var products Product

	err := r.db.Where("user_id =?", userID).Preload("ProductImages", "product_images.is_primary=1").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

//
