package produk

import (
	"github.com/jinzhu/gorm"
	//"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Produk, error)
	FindByUserID(userID int) ([]Produk, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Produk, error) {
	var produks []Produk
	err := r.db.Find(&produks).Error
	if err != nil {
		return produks, err
	}
	return produks, nil
}

func (r *repository) FindByUserID(userID int) ([]Produk, error) {
	var produks []Produk

	err := r.db.Where("user_id =?", userID).Find(&produks).Error
	if err != nil {
		return produks, err
	}
	return produks, nil
}
