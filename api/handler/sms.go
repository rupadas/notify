package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	common "github.com/rupadas/notify/config"
	"github.com/rupadas/notify/integration/sms"
)

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
