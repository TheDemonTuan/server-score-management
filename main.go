package main

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"os"
	"qldiemsv/common"
	"qldiemsv/router"
)

func init() {
	if os.Getenv("APP_ENV") == "development" {
		common.LoadEnvVar()
	}
	common.ConnectDB()
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
		ReduceMemoryUsage: true,
		Prefork:           os.Getenv("APP_ENV") == "production",
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
	if os.Getenv("APP_ENV") == "production" {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     os.Getenv("CLIENT_URL"),
			AllowCredentials: true,
		}))
	} else {
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOriginsFunc: func(origin string) bool {
				return os.Getenv("APP_ENV") == "development"
			},
		}))
	}
	app.Use(etag.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	//app.Use(cache.New(cache.Config{
	//	Expiration:   30 * time.Minute,
	//	CacheControl: true,
	//}))
	//app.Use(logger.New())

	router.SetupRouter(app)

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic("Error while starting server")
		return
	}
}
