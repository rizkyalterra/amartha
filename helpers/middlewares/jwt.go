package middlewares

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateTokenJWT(userId int32) (string, error) {
	// payload data
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate enkripsi + key
	return token.SignedString([]byte(viper.GetString(`jwt.Key`)))
}
