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

func AddEvent(c *fiber.Ctx) error {
	event := new(models.Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	environment, appId, err := helper.FetchEnvAndAppId(c)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON("Failed to fetch environment and appId")
	}
	event.Environment = environment
	event.AppId = appId
	database.DBConn.Create(&event)
	return c.Status(http.StatusOK).JSON(event)
}

type Channel struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

func AddEventChannels(c *fiber.Ctx) error {
	// Extract the array of objects from the request
	var channels []Channel
	if err := c.BodyParser(&channels); err != nil {
		return err
	}
	var eventId uint
	var err error
	var tempEventId uint64
	tempEventId, err = strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid event ID")
	}
	eventId = uint(tempEventId)
	environment, appId, err := helper.FetchEnvAndAppId(c)
	if err != nil {
		log.Println(err)
		log.Println(appId)
		return c.Status(http.StatusInternalServerError).JSON("Failed to fetch environment and appId")
	}
	var eventChannels []models.EventChannel
	for _, channel := range channels {
		eventChannel := models.EventChannel{}
		eventChannel.EventId = eventId
		eventChannel.ChannelId = channel.Id
		eventChannel.Environment = environment
		eventChannels = append(eventChannels, eventChannel)
	}
	database.DBConn.Create(&eventChannels)
	return c.Status(http.StatusOK).JSON(eventChannels)
}
