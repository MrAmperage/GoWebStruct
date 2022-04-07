package Modules

import "github.com/streadway/amqp"

type RabbitMQSetting struct {
	Connection *amqp.Connection
	ModuleType string
	Login      string
	Password   string
	Port       int64
	Adress     string
	Queues     []RabbitQueue
}

type RabbitQueue struct {
	Name  string
	Queue *amqp.Queue
}
