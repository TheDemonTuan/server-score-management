package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func studentRouter(r fiber.Router) {
	studentRouter := r.Group("students")

	studentRouter.Get("", controllers.StudentList)

	studentRouter.Get("/:id", controllers.StudentGetById)
	studentRouter.Post("", controllers.StudentCreate)
	studentRouter.Put("/:id", controllers.StudentUpdate)
	studentRouter.Delete("/:id", controllers.StudentDelete)
}
