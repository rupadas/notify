package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	database "github.com/rupadas/notify/config"
	"github.com/rupadas/notify/models"
)

func AddProvider(c *fiber.Ctx) error {
	provider := new(models.Provider)
	val, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		log.Println(ok)
		log.Println(val)
	}
	if err := c.BodyParser(provider); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	database.DBConn.Create(&provider)
	return c.Status(http.StatusOK).JSON(provider)
}

func AddProviderSettings(c *fiber.Ctx) error {
	providersetting := new(models.ProviderSetting)
	if err := c.BodyParser(providersetting); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var environment models.Environment
	val, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		log.Println(ok)
		log.Println(environment)
	}
	environment = val
	id, _ := strconv.Atoi(c.Params("id"))
	providersetting.ProviderId = uint(id)
	providersetting.Environment = environment
	log.Println(providersetting)
	database.DBConn.Create(&providersetting)
	return c.Status(200).JSON(providersetting)
}
