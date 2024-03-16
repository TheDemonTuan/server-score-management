package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func userRouter(r fiber.Router) {
	usersRoute := r.Group("users")

	usersRoute.Get("me", controllers.UserMe)
}
