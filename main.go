package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	Numbers []int32 `json:"numbers"`
}

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

	app.Post("/", func(c *fiber.Ctx) error {

		var body RequestBody
		var sum int32

		if err := c.BodyParser(&body); err != nil {
			log.Println("Error parsing body", err)

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		for _, v := range body.Numbers {
			sum += v
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"number": sum,
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		res := Response{
			Message: "oppsss you got lost in my simple app.... womp wompp",
		}

		return c.Status(http.StatusNotFound).JSON(res)
	})

	app.Listen(":3000")
}
