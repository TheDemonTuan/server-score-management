package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(a *fiber.App) {
	api := a.Group("api") // /api

	authRouter(api)
	departmentRouter(api)
}
