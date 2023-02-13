package routes

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	common "github.com/rupadas/raven/config"
	database "github.com/rupadas/raven/config"
	"github.com/rupadas/raven/integration/email"
	"github.com/rupadas/raven/integration/sms"
	"github.com/rupadas/raven/models"
)

func AddApp(c *fiber.Ctx) error {
	account := new(models.App)
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
	account.AccessKey = hex.EncodeToString(key)
	account.AccessToken = hex.EncodeToString((token))
	database.DBConn.Create(&account)
	return c.Status(http.StatusOK).JSON(account)
}

func AddCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	database.DBConn.Create(&customer)
	return c.Status(http.StatusOK).JSON(customer)
}

func AddChannel(c *fiber.Ctx) error {
	channel := new(models.Channel)
	val, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		log.Println(ok)
	}
	environment := val
	if err := c.BodyParser(channel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	channel.Environment = environment
	database.DBConn.Create(&channel)
	return c.Status(http.StatusOK).JSON(channel)
}

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

type ChannelObject struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

func AddEventChannels(c *fiber.Ctx) error {
	// Extract the array of objects from the request
	var channels []ChannelObject
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

type ProviderObject struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

func AddChannelProvider(c *fiber.Ctx) error {
	var providers []ProviderObject
	if err := c.BodyParser(&providers); err != nil {
		return err
	}
	var channelId uint
	var err error
	var tempChannelId uint64
	tempChannelId, err = strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid channel ID")
	}
	channelId = uint(tempChannelId)
	var environment models.Environment
	val, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		log.Println(ok)
		log.Println(environment)
	}
	environment = val
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

func AddProviderSetting(c *fiber.Ctx) error {
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

func SendEmail(c *fiber.Ctx) error {
	emailData := new(common.EmailBody)
	var appId uint
	val, ok := c.Locals("appId").(uint)
	if !ok {
		log.Println(ok)
	}
	appId = val
	log.Println("appId", appId)
	if err := c.BodyParser(emailData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	log.Println("emailData---", emailData)
	email.Send(*emailData)
	return c.Status(http.StatusOK).JSON(map[string]bool{"successfully sent email": true})
}

func SendSms(c *fiber.Ctx) error {
	smsData := new(common.SmsBody)
	var appId uint
	val, ok := c.Locals("appId").(uint)
	if !ok {
		log.Println(ok)
	}
	appId = val
	log.Println("appId", appId)
	if err := c.BodyParser(smsData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	log.Println("smsData---", smsData)
	sms.Send(*smsData)
	return c.Status(http.StatusOK).JSON(map[string]bool{"successfully sent sms": true})
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
