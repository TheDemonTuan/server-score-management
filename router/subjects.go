package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func subjectsRouter(r fiber.Router) {
	subjectsRoute := r.Group("subjects")

	subjectsRoute.Get("", controllers.SubjectGetList)
	subjectsRoute.Get("/:id", controllers.SubjectGetById)
	subjectsRoute.Post("", controllers.SubjectCreate)
	subjectsRoute.Put("/:id", controllers.SubjectUpdateById)
	//subjectsRoute.Delete("/:id", controllers.SubjectDeleteById)
	subjectsRoute.Put("/all", controllers.SubjectDeleteAll)
	subjectsRoute.Put("", controllers.SubjectDeleteList)
}
