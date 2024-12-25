package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

//go:generate mockgen -source=token.go -destination=mocks/token.go
type TokenManager interface {
	GenerateToken(fullname string, id int32) (string, error)
	ValidateToken(tokenString string) (*JwtCustomClaims, error)
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

	claims := &JwtCustomClaims{
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

func (m *Manager) ValidateToken(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return m.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, echo.ErrUnauthorized
	}

	return claims, nil
}
