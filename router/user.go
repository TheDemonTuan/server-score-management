package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func userRouter(r fiber.Router) {
	userRoute := r.Group("users")

	userRoute.Get("me", controllers.UserMe)
}
