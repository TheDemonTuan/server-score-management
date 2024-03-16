package controllers

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
	"qldiemsv/models/entity"
)

func UserMe(c *fiber.Ctx) error {
	userData, userDataIsOk := c.Locals("currentUserInfo").(entity.User)

	if !userDataIsOk {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		userData),
	)
}
