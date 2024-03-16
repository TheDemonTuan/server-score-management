package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func instructorsRouter(r fiber.Router) {
	instructorsRoute := r.Group("instructors")

	instructorsRoute.Get("", controllers.InstructorGetList)
	instructorsRoute.Get("/:id", controllers.InstructorGetById)
	instructorsRoute.Post("", controllers.InstructorCreate)
	instructorsRoute.Put("/:id", controllers.InstructorUpdateById)
	instructorsRoute.Delete("/:id", controllers.InstructorDeleteById)
}
