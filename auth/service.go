package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("CROWDFOUNDING_S3Cr3t_k3Y")

func NewService() *jwtService  {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}

func (s* jwtService) ValidateToken(encodeToken string) (*jwt.Token, error)  {
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if  !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}