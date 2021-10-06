package user

import (
	"net/http"

	"github.com/elangmfikri123/rest-api/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService Services
}

func NewHandler(userService Services) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) UserRegistration(c echo.Context) error {
	req := new(RequestUser)

	if err := c.Bind(req); err != nil {

		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, err := h.userService.CreateUser(*req)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	userData := UserResponseFormatter(newUser)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user succesfully register", userData)

	return c.JSON(http.StatusOK, response)
}
