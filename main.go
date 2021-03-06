package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elangmfikri123/rest-api/routes"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	fmt.Println("running...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Pre(middleware.RemoveTrailingSlash())

	routes.DefineApiRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
