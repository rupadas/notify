package sms

import (
	"log"

	common "github.com/rupadas/notify/config"
)

func Send(smsBody common.SmsBody) {
	log.Println(smsBody)
}
