package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func authRouter(r fiber.Router) {
	authRoute := r.Group("auth")

	authRoute.Post("login", controllers.AuthLogin)
	authRoute.Post("register", controllers.AuthRegister)
}
