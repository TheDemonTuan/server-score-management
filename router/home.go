package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func homeRouter(r fiber.Router) {
	// Group để định nghĩa bên main phải đi qua /api thì bên này chỉ x lý d liệu đi qua /api/home
	homeRoute := r.Group("home")

	// Định nghĩa các route con của /api/home
	homeRoute.Get("", controllers.HomeIndex)
}
