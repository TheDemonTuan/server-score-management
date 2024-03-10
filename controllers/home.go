package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// [GET] /
func HomeIndex(c *fiber.Ctx) error {
	return c.SendString("day la home dung get")
}
