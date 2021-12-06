package messaging

import "github.com/streadway/amqp"

type impl interface {
	NewConnect() error
	SendMessage(msg string, ch *amqp.Channel, nameQueue string) error
	CloseConn()

	// LOG QUEUE
	StartLogQueue() error
	GetLogQueueChannel() *amqp.Channel
	CloseLogQueue()
}

type service struct {
	impl
}

func NewService(messaging impl) service {
	return service{messaging}
}
