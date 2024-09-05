package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	Numbers []int32 `json:"number"`
}

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		body := c.Request().Body()
		sum := 0

		for _, v := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(string(body), "[", ""), "]", ""), ",") {
			num, err := strconv.Atoi(v)
			if err != nil {
				return c.Status(http.StatusOK).SendString("failed to parse")
			}
			sum += num
		}
		return c.Status(http.StatusOK).SendString(fmt.Sprintf("%d", sum))
	})

	app.Listen(":3000")
}
