package controllers

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
)

// [GET] /
func AuthLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(common.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Dang nhap thanh cong",
		Data:       nil,
	})
}
