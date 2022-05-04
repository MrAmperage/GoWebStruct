package RabbitMQModule

import (
	"encoding/json"
	"errors"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/streadway/amqp"
)

func (RabbitMQ *RabbitMQ) QueuesSubscribe() (Error error) {
	for Index, RabbitMQSubscribe := range RabbitMQ.RabbitMQChanel.Subscribes {
		RabbitMQ.RabbitMQChanel.Subscribes[Index].ChanelLink = RabbitMQ.RabbitMQChanel.Chanel
		RabbitMQ.RabbitMQChanel.Subscribes[Index].Messages, Error = RabbitMQ.RabbitMQChanel.Chanel.Consume(RabbitMQSubscribe.Queue, RabbitMQSubscribe.Consumer, RabbitMQSubscribe.AutoAck, RabbitMQSubscribe.Exclusive, RabbitMQSubscribe.noLocal, RabbitMQSubscribe.noWait, RabbitMQSubscribe.Args)
		if Error != nil {
			return Error
		}

	}
	return Error
}
func (RabbitMQChanel *RabbitMQChanel) GetSubscribeByQueueName(QueueName string) (RabbitMQSubscribe *RabbitMQSubscribe, Error error) {

	for _, Subscribe := range RabbitMQChanel.Subscribes {
		if Subscribe.Queue == QueueName {
			return &Subscribe, nil
		}

	}
	return RabbitMQSubscribe, errors.New("не найдена запрашиваемая очередь")

}
func (RabbitMQ *RabbitMQ) QueuesRiseAndBind() (Error error) {

	for _, QueueUP := range RabbitMQ.RabbitMQChanel.QueuesUP {

		QueueUP.Queue, Error = RabbitMQ.RabbitMQChanel.Chanel.QueueDeclare(QueueUP.Name, QueueUP.Durable,
			QueueUP.AutoDelete,
			QueueUP.Exclusive,
			QueueUP.NoWait,
			QueueUP.Args)
		if Error != nil {
			return Error
		}

		if len(QueueUP.Binding.Key) != 0 {
			Error = RabbitMQ.RabbitMQChanel.Chanel.QueueBind(QueueUP.Binding.Destination, QueueUP.Binding.Key, QueueUP.Binding.Source, QueueUP.Binding.NoWait, QueueUP.Binding.Args)
			if Error != nil {
				return Error
			}
		}

	}
	return Error
}

func (RabbitMQ *RabbitMQ) ExchangeRiseAndBind() (Error error) {
	for _, RabbitMQExchange := range RabbitMQ.RabbitMQChanel.ExchangeUP {
		Error := RabbitMQ.RabbitMQChanel.Chanel.ExchangeDeclare(RabbitMQExchange.ExchangeName, RabbitMQExchange.ExchangeType, RabbitMQExchange.Durable, RabbitMQExchange.AutoDelete, RabbitMQExchange.Internal, RabbitMQExchange.NoWait, RabbitMQExchange.Args)
		if Error != nil {

			return Error
		}
		if len(RabbitMQExchange.Binding.Key) != 0 {
			Error = RabbitMQ.RabbitMQChanel.Chanel.ExchangeBind(RabbitMQExchange.Binding.Destination, RabbitMQExchange.Binding.Key, RabbitMQExchange.Binding.Source, RabbitMQExchange.Binding.NoWait, RabbitMQExchange.Binding.Args)
			if Error != nil {
				return Error
			}
		}

	}
	return Error

}
func (RabbitMQSubscribe *RabbitMQSubscribe) GetMessageByCorrelationId(CorrelationId string) (RabbitMessage amqp.Delivery, Error error) {

	for Message := range RabbitMQSubscribe.Messages {
		if CorrelationId == Message.CorrelationId {

			if Message.Type == "Error" {
				Error = errors.New(string(Message.Body))

			}

			return Message, Error

		}

	}

	return RabbitMessage, errors.New("не найдено сообщение")
}
func (RabbitMQSubscribe *RabbitMQSubscribe) MessageProcessing(ORMs ...ORMModule.ORM) {

	for Message := range RabbitMQSubscribe.Messages {

		Function, HasFunction := RabbitMQSubscribe.MessageEmmiter.MessageHandlers[Message.Type+Message.RoutingKey]
		if HasFunction {
			Data, Error := Function(Message, ORMs)
			if Error != nil {

				RabbitMQSubscribe.ChanelLink.Publish("", Message.ReplyTo, false, false, amqp.Publishing{
					Type:          "Error",
					CorrelationId: Message.CorrelationId,
					Body:          []byte(Error.Error()),
				})

			} else {

				DataByte, _ := json.Marshal(Data)
				RabbitMQSubscribe.ChanelLink.Publish("", Message.ReplyTo, false, false, amqp.Publishing{
					Type:          "Answer",
					CorrelationId: Message.CorrelationId,
					Body:          DataByte,
				})

			}
		}

	}

}

func (MessageEmmiter *MessageEmmiter) Handler(RoutingKey string, EmmiterFunction EmmiterFunction) (Routing *RoutingObject) {
	return &RoutingObject{RoutingKey: RoutingKey, EmmiterFunction: EmmiterFunction, MessageEmmiterLink: MessageEmmiter}

}
func (RoutingObject *RoutingObject) Method(Method string) {
	if RoutingObject.MessageEmmiterLink.MessageHandlers == nil {
		RoutingObject.MessageEmmiterLink.MessageHandlers = make(map[string]EmmiterFunction)

	}
	RoutingObject.MessageType = Method
	HandlerKey := RoutingObject.MessageType + RoutingObject.RoutingKey
	RoutingObject.MessageEmmiterLink.MessageHandlers[HandlerKey] = RoutingObject.EmmiterFunction

}
