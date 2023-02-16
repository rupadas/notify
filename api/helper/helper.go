package helper

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rupadas/notify/models"
)

func FetchEnvAndAppId(c *fiber.Ctx) (models.Environment, uint, error) {
	env, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		return "", 0, errors.New("environment not found in locals")
	}

	appId, ok := c.Locals("appId").(uint)
	if !ok {
		return "", 0, errors.New("appid not found in locals")
	}
	return env, appId, nil
}
