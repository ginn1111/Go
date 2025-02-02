package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	SecretKey string
	Issuer    string
}

type jwtCustomClaims struct {
	Name  string `json:"Name"`
	Admin bool   `json:"Admin"`
	jwt.RegisteredClaims
}

func NewJWTService() JWTService {
	return &jwtService{
		SecretKey: getSecretKey(),
		Issuer:    "Gin",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "secret"
	}

	return secret
}

func (jwts *jwtService) GenerateToken(Name string, Admin bool) string {
	claims := &jwtCustomClaims{
		Name,
		Admin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jwts.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jwts.SecretKey))

	if err != nil {
		panic(err)
	}

	return ss
}

func (jwts *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method ")
		}

		return []byte(jwts.SecretKey), nil
	})
}
