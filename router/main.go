package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/middleware"
)

func SetupRouter(app *fiber.App) {
	//Các api public mà không cần phải đăng nhập để truy cập
	publicAPI := app.Group("api")
	authRouter(publicAPI)

	//Các api cần phải đăng nhập để truy cập
	privateAPI := app.Group("api", middleware.Protected)
	departmentRouter(privateAPI)
	subjectRouter(privateAPI)
}
