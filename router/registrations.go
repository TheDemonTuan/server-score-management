package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func registrationsRouter(r fiber.Router) {
	registrationsRoute := r.Group("registrations")

	registrationsRoute.Add("GET", "", controllers.InstructorGetAll)
}
