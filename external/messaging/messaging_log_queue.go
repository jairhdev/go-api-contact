package messaging

import "github.com/streadway/amqp"

const LogQueueName string = "log_queue"

var logQueueChannel *amqp.Channel

// **********************************

func (data mq) StartLogQueue() error {
	var err error

	// channel
	if logQueueChannel == nil {
		logQueueChannel, err = conn.Channel()
		if err != nil {
			return err
		}
	}

	// queue
	_, err = logQueueChannel.QueueDeclare(
		LogQueueName, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	return err
}

func (data mq) GetLogQueueChannel() *amqp.Channel {
	return logQueueChannel
}

func (data mq) CloseLogQueue() {
	logQueueChannel.Close()
}
