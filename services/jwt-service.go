package services

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type JwtService interface {
	GenerateToken(userId uint64) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

// get secret key
func (s *jwtService) SecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey != "" {
		secretKey = "secret"
	}
	return secretKey
}

// generate token
func (s *jwtService) GenerateToken(userId uint64) string {
	claim := jwtCustomClaim{userId, jwt.StandardClaims{}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, _ := token.SignedString([]byte(s.SecretKey()))
	return signedToken
}

// validate token
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey()), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
