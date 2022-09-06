package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
)


var secret string = os.Getenv("JWT_SECRET")

func GenerateToken(uid uint32) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"]=uid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	 t, err := token.SignedString([]byte(secret))
	 if err!=nil{
		return "",err
	 }
	return t, err
}