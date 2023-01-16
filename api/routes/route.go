package routes

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	database "github.com/rupadas/raven/config"
	"github.com/rupadas/raven/models"
)

func AddApp(c *fiber.Ctx) error {
	account := new(models.AppModel)
	if err := c.BodyParser(account); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// Generate a 32-byte key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	// Generate a 64-byte token
	token := make([]byte, 64)
	_, err = rand.Read(token)
	if err != nil {
		panic(err)
	}
	account.Key = hex.EncodeToString(key)
	account.Token = hex.EncodeToString((token))
	database.DBConn.Create(&account)
	return c.Status(http.StatusOK).JSON(account)
}

func AddCustomer(c *fiber.Ctx) error {
	customer := new(models.CustomerModel)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	database.DBConn.Create(&customer)
	return c.Status(http.StatusOK).JSON(customer)
}

func AddChannel(c *fiber.Ctx) error {
	channel := new(models.ChannelModel)
	if err := c.BodyParser(channel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	database.DBConn.Create(&channel)
	return c.Status(http.StatusOK).JSON(channel)
}

func AddEvent(c *fiber.Ctx) error {
	event := new(models.EventModel)
	if err := c.BodyParser(event); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
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

	var eventChannels []models.EventChannelModel
	for _, channel := range channels {
		eventChannel := models.EventChannelModel{}
		eventChannel.EventId = eventId
		eventChannel.ChannelId = channel.Id
		eventChannel.Environment = "DEVELOPMENT"
		eventChannels = append(eventChannels, eventChannel)
	}
	database.DBConn.Create(&eventChannels)
	return c.Status(http.StatusOK).JSON(eventChannels)
}
