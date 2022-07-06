package commands

import (
	"net/http"

	"github.com/alash3al/scraply/pkg/fetch"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/urfave/cli/v2"
)

// HTTPServer a factory that creates a http server
func HTTPServer() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		app := fiber.New()

		app.Use(recover.New())

		if ctx.Bool("logging") {
			app.Use(logger.New())
		}

		app.Use(compress.New(compress.Config{
			Level: compress.LevelBestCompression,
		}))

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		app.Post("/extract", func(c *fiber.Ctx) error {
			var input fetch.Input

			if err := c.BodyParser(&input); err != nil {
				return c.Status(422).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			output, err := fetch.Do(c.Context(), fetch.Input{
				URL:        input.URL,
				Method:     http.MethodGet,
				Body:       nil,
				UserAgent:  input.UserAgent,
				ReturnBody: input.ReturnBody,
				Extractors: input.Extractors,
			})
			if err != nil {
				return c.Status(417).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			return c.Status(200).JSON(output)
		})

		return app.Listen(ctx.String("address"))
	}
}
