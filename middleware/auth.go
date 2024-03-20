package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func Protected(c *fiber.Ctx) error {
	currUserId, err := func() (string, error) {
		jwtToken := c.Cookies(os.Getenv("JWT_NAME"), "")

		if jwtToken == "" {
			return "", errors.New("JWT not found")
		}

		token, tokenIsErr := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, isOk := token.Method.(*jwt.SigningMethodHMAC); !isOk {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if tokenIsErr != nil || !token.Valid {
			return "", errors.New("Invalid or expired JWT")
		}

		tokenClaims, tokenClaimsIsOk := token.Claims.(jwt.MapClaims)

		if !tokenClaimsIsOk {
			return "", errors.New("Invalid JWT payload")
		}

		userId, userIdIsOk := tokenClaims["uid"].(string)

		if !userIdIsOk {
			return "", errors.New("Invalid JWT payload")
		}

		return userId, nil
	}()

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	if currUserId == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid JWT payload")
	}

	c.Locals("currUserId", currUserId)
	return c.Next()
}
