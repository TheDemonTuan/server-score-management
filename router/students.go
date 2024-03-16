package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func studentsRouter(r fiber.Router) {
	studentsRoute := r.Group("students")

	studentsRoute.Get("", controllers.StudentGetList)
	studentsRoute.Get("/:id", controllers.StudentGetById)
	studentsRoute.Post("", controllers.StudentCreate)
	studentsRoute.Put("/:id", controllers.StudentUpdateById)
	studentsRoute.Delete("/:id", controllers.StudentDeleteById)
}
