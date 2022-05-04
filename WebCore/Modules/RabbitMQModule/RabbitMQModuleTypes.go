package RabbitMQModule

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
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
	Binding    RabbitMQBinding
}

type RabbitMQSubscribe struct {
	Messages       <-chan amqp.Delivery
	Queue          string
	Consumer       string
	AutoAck        bool
	Exclusive      bool
	noLocal        bool
	noWait         bool
	Args           amqp.Table
	ChanelLink     *amqp.Channel
	MessageEmmiter MessageEmmiter
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
	Binding      RabbitMQBinding
}

type RabbitMQBinding struct {
	Destination string
	Key         string
	Source      string
	NoWait      bool
	Args        amqp.Table
}
type RabbitMQChanel struct {
	Chanel     *amqp.Channel
	QueuesUP   []RabbitMQQueue
	Subscribes []RabbitMQSubscribe
	ExchangeUP []RabbitMQExchange
}
type RabbitMQMessageCallbackFunction func(RabbitMQMessage amqp.Delivery)

type MessageEmmiter struct {
	MessageHandlers map[string]EmmiterFunction
}
type RoutingObject struct {
	MessageType        string
	RoutingKey         string
	EmmiterFunction    EmmiterFunction
	MessageEmmiterLink *MessageEmmiter
}

type EmmiterFunction func(Message amqp.Delivery, ORMs []ORMModule.ORMInterface) (Data any, Error error)
