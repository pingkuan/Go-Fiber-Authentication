package middlewares

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
)

func Protect(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return errors.New("authorization header is required")
	}

    var Secret string = os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(
		authHeader[7:],
		func(token *jwt.Token) (interface{}, error) {
			
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(Secret), nil
		},
	)

	if err != nil {
		return errors.New("error parsing token")
	}

	claims, ok :=token.Claims.(jwt.MapClaims)

	if !(ok && token.Valid){
		return errors.New("invalid token")
	}

	if expiresAt, ok := claims["exp"]; ok&& int64(expiresAt.(float64))<time.Now().UTC().Unix(){
		return errors.New("token expired")
	}

	c.Locals("ID", claims["user_id"])

	return c.Next()		   
}