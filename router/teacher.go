package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func teacherRouter(r fiber.Router) {
	teacherRouter := r.Group("teacher")

	teacherRouter.Get("", controllers.TeacherList)
	teacherRouter.Get("/:id", controllers.TeacherGetById)
	teacherRouter.Post("", controllers.TeacherCreate)
	teacherRouter.Put("/:id", controllers.TeacherUpdate)
	teacherRouter.Delete("/:id", controllers.TeacherDelete)
}
