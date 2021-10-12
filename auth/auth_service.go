package auth

import (
	"os"
	"time"

	"github.com/elangmfikri123/rest-api/middleware"
	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	GetAccessToken(userID uint) (string, error)
}

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) GetAccessToken(userID uint) (string, error) {
	claims := &middleware.JwtCustomClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedkey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedkey, err
	}

	return signedkey, nil
}
