package messaging

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection

// **********************************

func (data mq) NewConnect() error {
	var err error

	if conn == nil {
		url := fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			os.Getenv("MQ_RABBIT_USER"),
			os.Getenv("MQ_RABBIT_PWD"),
			os.Getenv("MQ_RABBIT_HOST"),
			os.Getenv("MQ_RABBIT_PORT"),
		)

		conn, err = amqp.Dial(url)
	}
	return err
}

func (data mq) SendMessage(msg string, ch *amqp.Channel, nameQueue string) error {
	return ch.Publish(
		"",        // exchange
		nameQueue, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

func (data mq) CloseConn() {
	conn.Close()
}
