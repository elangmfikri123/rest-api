package user

import (
	"github.com/elangmfikri123/rest-api/database"
	"github.com/elangmfikri123/rest-api/helper"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	db := database.GetDbInstance()
	db.AutoMigrate(User{})
	userRepo := NewRepository(db)
	userService := NewServices(userRepo)
	userHandler := NewHandler(userService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.UserRegistration,
		},
	}
}
