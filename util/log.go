package util

import (
	"log"

	"github.com/jairhdev/go-api-contact/external/messaging"
)

func NewLog(msg string) {
	var mq = messaging.NewService(messaging.NewMessaging())

	if err := mq.StartLogQueue(); err != nil {
		log.Printf("error newlog().startlogqueue(): %v", err)
	}

	if err := mq.SendMessage(msg, mq.GetLogQueueChannel(), messaging.LogQueueName); err != nil {
		log.Printf("error newlog().sendmessage(): %v", err)
	}
}
