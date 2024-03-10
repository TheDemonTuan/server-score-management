package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// [GET] /
func AuthLogin(c *fiber.Ctx) error {
	return c.SendString("dang nhap")
}
