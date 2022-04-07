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
