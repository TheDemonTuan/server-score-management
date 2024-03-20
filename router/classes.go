package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func classesRouter(r fiber.Router) {
	classesRoute := r.Group("classes")

	classesRoute.Get("", controllers.ClassGetList)
	classesRoute.Get("/:id", controllers.ClassGetById)
	classesRoute.Post("", controllers.ClassCreate)
	classesRoute.Put("/:id", controllers.ClassUpdate)
	//classesRoute.Delete("/:id", controllers.ClassDelete)
	classesRoute.Delete("/all", controllers.ClassDeleteAll)
	classesRoute.Delete("", controllers.ClassDeleteList)
}
