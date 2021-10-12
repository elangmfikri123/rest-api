package user

import (
	"net/http"

	"github.com/elangmfikri123/rest-api/auth"
	"github.com/elangmfikri123/rest-api/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService Services
	authService auth.AuthService
}

func NewHandler(userService Services, authService auth.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) UserRegistration(c echo.Context) error {
	req := new(RequestUser)

	if err := c.Bind(req); err != nil {

		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Permintaan Tidak Valid", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "ERROR", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	existEmail := h.userService.CheckExistEmail(*req)
	if existEmail != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "ERROR", existEmail.Error(), nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, err := h.userService.CreateUser(*req)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "ERROR", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	auth_token, err := h.authService.GetAccessToken(newUser.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "ERROR", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := UserResponseFormatter(newUser, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "SUKSES", "Pengguna Berhasil Mendaftar", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) UserLogin(c echo.Context) error {
	req := new(RequestUserLogin)

	if err := c.Bind(req); err != nil {

		response := helper.ResponseFormatter(http.StatusBadRequest, "ERROR", "Permintaan Tidak Valid", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "ERROR", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	userAuth, err := h.userService.AuthUser(*req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "ERROR", err.Error(), nil)

		return c.JSON(http.StatusUnauthorized, response)
	}

	auth_token, err := h.authService.GetAccessToken(userAuth.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "ERROR", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := UserResponseFormatter(userAuth, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "SUKSES", "Pengguna Diontetikasi", userData)

	return c.JSON(http.StatusOK, response)

}
