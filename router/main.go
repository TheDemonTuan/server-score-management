package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/middleware"
)

func SetupRouter(app *fiber.App) {
	publicAPIRoute := app.Group("api")
	authRouter(publicAPIRoute)

	privateAPIRoute := app.Group("api", middleware.Protected)
	usersRouter(privateAPIRoute)
	departmentsRouter(privateAPIRoute)
	subjectsRouter(privateAPIRoute)
	instructorsRouter(privateAPIRoute)
	studentsRouter(privateAPIRoute)
}
