package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/request"
)

func AuthLogin(c *fiber.Ctx) error {
	body := new(request.AuthLogin)

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	return c.Status(fiber.StatusOK).JSON(common.NewResponse(fiber.StatusOK, "Login success", body))
}

func AuthRegister(c *fiber.Ctx) error {
	bodyData, err := common.Validator[request.AuthRegister](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var userRecord entity.User

	if err := common.DBConn.First(&userRecord, "user_name = ?", bodyData.UserName).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusInternalServerError, "Error while fetching user")
		}
	}

	if userRecord != (entity.User{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Username already exists")
	}

	hashPassword, hashPasswordErr := bcrypt.GenerateFromPassword([]byte(bodyData.Password), 11)

	if hashPasswordErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error while hashing password")
	}

	newUser := entity.User{
		FirstName: bodyData.FirstName,
		LastName:  bodyData.LastName,
		UserName:  bodyData.UserName,
		Password:  string(hashPassword),
	}

	if err := common.DBConn.Create(&newUser).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error while creating new user")
	}

	return c.Status(fiber.StatusOK).JSON(common.NewResponse(fiber.StatusOK, "Register success", nil))
}
