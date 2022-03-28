package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guionardo/auth-service/golang/setup"
)

func checkFeedAPIKey(c *fiber.Ctx) error {
	feedApiKey := c.Get("FEED_API_KEY", "")
	if len(feedApiKey) == 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "MISSING FEED_API_KEY HEADER")
	}
	cfg := setup.GetConfiguration()

	if cfg.FEED_API_KEY != feedApiKey {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "FEED_API_KEY HEADER IS INVALID")
	}

	return nil
}
