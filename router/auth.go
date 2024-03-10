package router

import (
	"qldiemsv/controllers"

	"github.com/gofiber/fiber/v2"
)

func authRouter(r fiber.Router) {
	homeRoute := r.Group("auth")

	homeRoute.Get("", controllers.AuthLogin)
}
