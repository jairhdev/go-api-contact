package messaging

import (
	"testing"

	"github.com/jairhdev/go-api-contact/config"
)

func TestService(t *testing.T) {

	//****************************
	// Start config TEST
	const env string = "test"
	if err := config.NewConfig(env); err != nil {
		panic(err)
	}
	//****************************

	var mq = NewService(NewMessaging())

	t.Run("testa se conexão é realizada com sucesso", func(t *testing.T) {
		err := mq.NewConnect()

		if err != nil {
			t.Errorf("\nExpected= %v\nResult= %v\n", nil, err)
		}
	})

	t.Run("testa envio de mensagem", func(t *testing.T) {
		const msg string = "test message. please ignore."

		if err := mq.StartLogQueue(); err != nil {
			t.Errorf("\nExpected= %v\nResult= %v\n", nil, err)
		}

		if err := mq.SendMessage(msg, mq.GetLogQueueChannel(), LogQueueName); err != nil {
			t.Errorf("\nExpected= %v\nResult= %v\n", nil, err)
		}
	})
}
