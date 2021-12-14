package product

import (
	"github.com/elangmfikri123/rest-api/database"
	"github.com/elangmfikri123/rest-api/helper"
	"github.com/labstack/echo/v4"
)

// var db *gorm.DB

type ProductRoutes struct {
}

func (r ProductRoutes) Route() []helper.Route {
	db := database.GetDbInstance()
	db.AutoMigrate(Product{})
	productRepository := NewRepository(db)
	// productService := NewServices(productRepository)

	// productHandler := NewHandler(productService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/insertproduct",
			Handler: productRepository.ProductCreate,
		},
	}
}

func NewServices(productRepo *repository) {
	panic("unimplemented")
}
