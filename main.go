package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		res := Response{
			Message: "Hi welcome to my simple app!",
		}
		return c.Status(http.StatusOK).JSON(res)
	})

	app.Use(func(c *fiber.Ctx) error {
		res := Response{
			Message: "oppsss you got lost in my simple app.... womp wompp",
		}

		return c.Status(http.StatusNotFound).JSON(res)
	})

	app.Listen(":3000")
}
