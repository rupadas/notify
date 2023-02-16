package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	database "github.com/rupadas/notify/config"
	"github.com/rupadas/notify/helper"
	"github.com/rupadas/notify/models"
)

func AddChannel(c *fiber.Ctx) error {
	channel := new(models.Channel)
	if err := c.BodyParser(channel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	environment, appId, err := helper.FetchEnvAndAppId(c)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON("Failed to fetch environment and appId")
	}
	channel.Environment = environment
	channel.AppId = appId
	database.DBConn.Create(&channel)
	return c.Status(http.StatusOK).JSON(channel)
}

type Provider struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

func AddChannelProviders(c *fiber.Ctx) error {
	var providers []Provider
	if err := c.BodyParser(&providers); err != nil {
		return err
	}
	environment, appId, err := helper.FetchEnvAndAppId(c)
	if err != nil {
		log.Println(err)
		log.Println(appId)
		return c.Status(http.StatusInternalServerError).JSON("Failed to fetch environment and appId")
	}
	var channelId uint
	var tempChannelId uint64
	tempChannelId, err = strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid channel ID")
	}
	channelId = uint(tempChannelId)
	var providerChannels []models.ChannelProvider
	for _, provider := range providers {
		providerChannel := models.ChannelProvider{}
		providerChannel.ProviderId = provider.Id
		providerChannel.ChannelId = channelId
		providerChannel.Environment = environment
		providerChannels = append(providerChannels, providerChannel)
	}
	database.DBConn.Create(&providerChannels)
	return c.Status(http.StatusOK).JSON(providerChannels)
}

func FetchEventChannles(c *fiber.Ctx) error {
	var eventId uint64
	var err error
	eventId, err = strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid event ID")
	}
	var event models.Event
	if err := database.DBConn.Preload("Channels").First(&event, eventId).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.JSON(event)
}
