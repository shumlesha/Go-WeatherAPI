package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	jwt.StandardClaims
}

type JwtUtil interface {
	GenerateToken(userId uuid.UUID, username string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type jwtUtil struct {
	secretKey string
}

func NewJwtUtil(secretKey string) JwtUtil {
	return &jwtUtil{
		secretKey: secretKey,
	}
}

func (j jwtUtil) GenerateToken(userId uuid.UUID, username string) (string, error) {
	claims := &Claims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j jwtUtil) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrSignatureInvalid
	}
}
