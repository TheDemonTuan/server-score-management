package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(a *fiber.App) {
	// Group để định nghĩa việc xử lý các request phải đi qua /api
	api := a.Group("api") // /api

	// Định nghĩa các route con của /api
	// /api/home
	departmentRouter(api)

}
