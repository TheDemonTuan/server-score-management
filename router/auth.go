package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
	"qldiemsv/middleware"
)

func authRouter(r fiber.Router) {
	authRoute := r.Group("auth")

	authRoute.Post("login", controllers.AuthLogin)
	authRoute.Post("register", controllers.AuthRegister)
	authRoute.Get("verify", middleware.Protected, controllers.AuthVerify)
	authRoute.Delete("logout", middleware.Protected, controllers.AuthLogout)
}
