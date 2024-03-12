package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func subjectRouter(r fiber.Router) {
	subjectRouter := r.Group("subject")

	subjectRouter.Get("", controllers.SubjectList)

	subjectRouter.Get("/:id", controllers.SubjectGetById)
	subjectRouter.Post("", controllers.SubjectCreate)
	subjectRouter.Put("/:id", controllers.SubjectUpdate)
	subjectRouter.Delete("/:id", controllers.SubjectDelete)
}
