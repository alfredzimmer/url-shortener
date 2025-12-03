package routes

import (
	"github.com/alfredzimmer/url-shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in the database"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to db",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}

// In production this end point should be protected.
func ResolveRateLimit(c *fiber.Ctx) error {
	r2 := database.CreateClient(1)
	defer r2.Close()

	err := r2.Del(database.Ctx, c.IP()).Err()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not find IP",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Rate limit reset successfully",
		"ip":      c.IP(),
	})
}
