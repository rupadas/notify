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

func AddProvider(c *fiber.Ctx) error {
	provider := new(models.Provider)
	val, ok := c.Locals("Environment").(models.Environment)
	if !ok {
		log.Println(ok)
	}
	environment := val
	provider.Environment = environment
	if err := c.BodyParser(provider); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	database.DBConn.Create(&provider)
	return c.Status(http.StatusOK).JSON(provider)
}

type provider struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

type ProviderRule struct {
	Country   string `json:"name"`
	providers []provider
}

func AddChannelProvider(c *fiber.Ctx) error {
	var ProviderRule ProviderRule
	if err := c.BodyParser(&ProviderRule); err != nil {
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
	}
	environment = val

	var providerChannels []models.ChannelProviderRule
	for _, provider := range ProviderRule.providers {
		providerChannel := models.ChannelProviderRule{}
		providerChannel.ProviderId = provider.Id
		providerChannel.ChannelId = channelId
		providerChannel.Environment = environment
		providerChannels = append(providerChannels, providerChannel)
	}
	database.DBConn.Create(&providerChannels)
	return c.Status(http.StatusOK).JSON(providerChannels)
}

func UpdateProvider(c *fiber.Ctx) error {
	provider := new(models.Provider)
	if err := c.BodyParser(provider); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	id, _ := strconv.Atoi(c.Params("id"))
	database.DBConn.Model(&models.Provider{}).Where("id = ?", id).Updates(models.Provider{AccessKey: provider.AccessKey, AccessToken: provider.AccessToken})
	return c.Status(200).JSON(provider)
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
