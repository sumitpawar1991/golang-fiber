package controller

import "github.com/gofiber/fiber/v3"

func BlogList(c fiber.Ctx) error {

	records := []any{}

	return c.JSON(fiber.Map{
		"status":       "ok",
		"message":      "blog list",
		"blog_records": records,
	})
}
