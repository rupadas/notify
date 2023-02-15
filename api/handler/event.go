package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	database "github.com/rupadas/notify/config"
	"github.com/rupadas/notify/models"
)

func AddEvent(c *fiber.Ctx) error {
	event := new(models.Event)
	val, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		log.Println(ok)
	}
	environment := val
	if err := c.BodyParser(event); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	event.Environment = environment
	log.Println(event)
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
	var environment models.Environment
	tempEventId, err = strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid event ID")
	}
	eventId = uint(tempEventId)
	val, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		log.Println(ok)
	}
	environment = val
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
