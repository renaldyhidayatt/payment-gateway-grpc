package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type TokenManager interface {
	GenerateToken(fullname string, id int32) (string, error)
	ValidateToken(tokenString string) (*jwtCustomClaims, error)
}

type Manager struct {
	secretKey []byte
}

func NewManager(secretKey string) (*Manager, error) {
	if secretKey == "" {
		return nil, errors.New("empty secret key")
	}
	return &Manager{secretKey: []byte(secretKey)}, nil
}

func (m *Manager) GenerateToken(fullname string, id int32) (string, error) {

	claims := &jwtCustomClaims{
		fullname,
		true,
		jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secretKey)
}

func (m *Manager) ValidateToken(tokenString string) (*jwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return m.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok || !token.Valid {
		return nil, echo.ErrUnauthorized
	}

	return claims, nil
}
