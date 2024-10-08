package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func JWTgenerate(username string) (string, error) {
	claims := jwt.MapClaims{
		"jti":    uuid.New().String(),
		"aud":    "3rd-party-auth.com",
		"iss":    "3rd-party-auth.com",
		"iat":    time.Now().Unix(),
		"user":   username,
		"expire": time.Now().Add(time.Minute * 60).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_S_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_S_KEY")),
		ErrorHandler: jwtErrr,
	})
}

func jwtErrr(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "unauthorized request",
			"message": "user not authorized to proced",
		})
	}
	return nil
}
