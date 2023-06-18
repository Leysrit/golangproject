package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("secret")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"userID": userID,
	// 	"exp": time.Now().Add(time.Hour * 24).Unix(),
	// })

	// tokenString, err := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	return "", err
	// }
	// return tokenString, nil

	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	
	return token, nil
}
