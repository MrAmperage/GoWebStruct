package Modules

import "github.com/streadway/amqp"

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
type RabbitMQChanel struct {
	Chanel     *amqp.Channel
	QueuesUP   []RabbitMQQueue
	Subscribes []RabbitMQSubscribe
}

func (RabbitMQSetting *RabbitMQ) QueuesSubscribe() (Error error) {
	for _, RabbitMQSubscribe := range RabbitMQSetting.RabbitMQChanel.Subscribes {
		RabbitMQSubscribe.Messages, Error = RabbitMQSetting.RabbitMQChanel.Chanel.Consume(RabbitMQSubscribe.Name, "", true, false, false, false, nil)
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
