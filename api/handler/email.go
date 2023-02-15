package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	common "github.com/rupadas/notify/config"
	"github.com/rupadas/notify/integration/email"
)

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
