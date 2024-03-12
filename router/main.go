package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/middleware"
)

func SetupRouter(app *fiber.App) {
	publicAPI := app.Group("api")
	authRouter(publicAPI)

	privateAPI := app.Group("api", middleware.Protected)
	userRouter(privateAPI)
	departmentRouter(privateAPI)
	subjectRouter(privateAPI)
}
