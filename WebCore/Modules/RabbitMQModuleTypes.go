package Modules

import "github.com/streadway/amqp"

type RabbitMQSetting struct {
	Connection     *amqp.Connection
	ModuleType     string
	Login          string
	Password       string
	Port           int64
	Adress         string
	RabbitMQChanel RabbitMQChanel
}

type RabbitMQQueue struct {
	Name  string
	Queue amqp.Queue
}

type RabbitMQSubscribe struct {
	Name     string
	Messages <-chan amqp.Delivery
}
type RabbitMQChanel struct {
	Chanel     *amqp.Channel
	QueuesUP   []RabbitMQQueue
	Subscribes []RabbitMQSubscribe
}

func (RabbitMQSetting *RabbitMQSetting) QueuesSubscribe() (Error error) {
	for _, RabbitMQSubscribe := range RabbitMQSetting.RabbitMQChanel.Subscribes {
		RabbitMQSubscribe.Messages, Error = RabbitMQSetting.RabbitMQChanel.Chanel.Consume(RabbitMQSubscribe.Name, "", true, false, false, false, nil)
		if Error != nil {
			return Error
		}

	}
	return Error
}
