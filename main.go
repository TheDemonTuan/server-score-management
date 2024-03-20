package main

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
		ReduceMemoryUsage: false,
		Prefork:           true,
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

	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CLIENT_URL"),
		AllowCredentials: true,
	}))
	app.Use(etag.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	if os.Getenv("APP_ENV") == "development" {
		app.Use(logger.New())
	}
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    os.Getenv("COOKIE_SECRET"),
		Except: []string{os.Getenv("JWT_NAME")},
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Quan Ly Diem Sinh Vien Metrics"}))

	router.SetupRouter(app)

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
