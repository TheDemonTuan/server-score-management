package main

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"qldiemsv/common"
	"qldiemsv/router"
)

func init() {
	common.LoadEnvVar()
	common.ConnectDB()
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		ReduceMemoryUsage: true,
		Prefork:           false,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Quan Ly Diem Sinh Vien",
		AppName:           "Quan Ly Diem Sinh Vien v1.0.0",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(code).JSON(common.NewResponse(
				code,
				err.Error(),
				nil))
		},
	})

	//Testing
	//app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CLIENT_URL"),
		AllowCredentials: true,
	}))
	app.Use(etag.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(logger.New())

	router.SetupRouter(app)

	err := app.Listen(os.Getenv("PORT"))
	if err != nil {
		panic("Error while starting server")
		return
	}
}
