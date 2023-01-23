package sms

import (
	"log"

	common "github.com/rupadas/raven/config"
)

func Send(smsBody common.SmsBody) {
	log.Println(smsBody)
}
