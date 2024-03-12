package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"os"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"strings"
)

func Protected(c *fiber.Ctx) error {
	currentUserInfoData, err := func() (entity.User, error) {
		jwtToken := c.Cookies(os.Getenv("JWT_NAME"), "")
		if jwtToken == "" {
			headerToken := c.Get("Authorization")
			if headerToken == "" {
				return entity.User{}, errors.New("Missing token")
			}

			tokenFields := strings.Fields(headerToken)

			if len(tokenFields) != 2 || strings.ToLower(tokenFields[0]) != "bearer" || tokenFields[1] == "" {
				return entity.User{}, errors.New("Bad authorization header")
			}

			jwtToken = tokenFields[1]
		}

		token, tokenIsErr := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, isOk := token.Method.(*jwt.SigningMethodHMAC); !isOk {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if tokenIsErr != nil || !token.Valid {
			return entity.User{}, errors.New("Invalid or expired JWT")
		}

		tokenClaims, tokenClaimsIsOk := token.Claims.(jwt.MapClaims)

		if !tokenClaimsIsOk {
			return entity.User{}, errors.New("Invalid JWT payload")
		}

		userId, userIdIsOk := tokenClaims["uid"].(string)

		if !userIdIsOk {
			return entity.User{}, errors.New("Invalid JWT payload")
		}

		var userRecord entity.User
		if err := common.DBConn.First(&userRecord, "id = ?", userId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return entity.User{}, errors.New("User not found")
			}
			return entity.User{}, errors.New("Error while fetching user")
		}

		return userRecord, nil
	}()

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	if currentUserInfoData == (entity.User{}) {
		return fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	c.Locals("currentUserInfo", currentUserInfoData)

	return c.Next()
}
