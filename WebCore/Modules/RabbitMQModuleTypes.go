package Modules

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection     *amqp.Connection
	ModuleType     string
	Login          string
	Password       string
	Port           int64
	Adress         string
	RabbitMQChanel RabbitMQChanel
}

type RabbitMQQueue struct {
	Name       string
	Queue      amqp.Queue
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type RabbitMQSubscribe struct {
	Name     string
	Messages <-chan amqp.Delivery
}
type RabbitMQExchange struct {
	ExchangeName string
	QueueName    string
	ExchangeType string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Args         amqp.Table
}
type RabbitMQChanel struct {
	Chanel     *amqp.Channel
	QueuesUP   []RabbitMQQueue
	Subscribes []RabbitMQSubscribe
	ExchangeUP []RabbitMQExchange
}

func (RabbitMQ *RabbitMQ) QueuesSubscribe() (Error error) {
	for _, RabbitMQSubscribe := range RabbitMQ.RabbitMQChanel.Subscribes {
		RabbitMQSubscribe.Messages, Error = RabbitMQ.RabbitMQChanel.Chanel.Consume(RabbitMQSubscribe.Name, "", true, false, false, false, nil)
		if Error != nil {
			return Error
		}

	}
	return Error
}

func (RabbitMQ *RabbitMQ) QueuesRise() (Error error) {

	for _, QueueUP := range RabbitMQ.RabbitMQChanel.QueuesUP {

		QueueUP.Queue, Error = RabbitMQ.RabbitMQChanel.Chanel.QueueDeclare(QueueUP.Name, QueueUP.Durable,
			QueueUP.AutoDelete,
			QueueUP.Exclusive,
			QueueUP.NoWait,
			QueueUP.Args)
		if Error != nil {
			return Error
		}

	}
	return Error
}

func (RabbitMQ *RabbitMQ) ExchangeRise() (Error error) {
	for _, RabbitMQExchange := range RabbitMQ.RabbitMQChanel.ExchangeUP {
		Error := RabbitMQ.RabbitMQChanel.Chanel.ExchangeDeclare(RabbitMQExchange.ExchangeName, RabbitMQExchange.ExchangeType, RabbitMQExchange.Durable, RabbitMQExchange.AutoDelete, RabbitMQExchange.Internal, RabbitMQExchange.NoWait, RabbitMQExchange.Args)
		if Error != nil {

			return Error
		}
	}
	return Error

}
