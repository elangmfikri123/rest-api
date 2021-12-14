package produk

import (
	"github.com/jinzhu/gorm"
	//"gorm.io/gorm"
)

type Repository interface {
	//FindAll() ([]Produk, error)
	FindAll(id int) *Produk
	FindByUserID(id int) (Produk, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(id string) *Produk {
	var produks Produk
	err := r.db.Preload("ProdukImages", "produk_images.is_primary = 1").Find(&produks).Error
	if err == nil {
		return &produks
	}

	return nil
}

func (r *repository) FindByUserID(userID int) (Produk, error) {
	var produks Produk

	err := r.db.Where("user_id =?", userID).Preload("ProdukImages", "produk_images.is_primary=1").Find(&produks).Error
	if err != nil {
		return produks, err
	}
	return produks, nil
}

//
