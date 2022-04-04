package consumer

// this package is used for receiving notification from the blockchain service

import (
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
	amqp "github.com/wagslane/go-rabbitmq"
)

type Consumer struct {
	mqConsumer amqp.Consumer
}

const reconnectInterval = time.Second * 10

func NewConsumer(uri string) (c *Consumer, err error) {
	c = &Consumer{}
	c.mqConsumer, err = amqp.NewConsumer(
		uri,
		amqp091.Config{},
		amqp.WithConsumerOptionsLogging,
		amqp.WithConsumerOptionsReconnectInterval(reconnectInterval),
	)
	return
}

func (c Consumer) StartConsuming() (err error) {
	err = c.mqConsumer.StartConsuming(
		updateAccountStatus,
		"exchange-topic", // TODO: give this a name
		[]string{},
		// TODO: variablize this
		amqp.WithConsumeOptionsBindingExchangeName("exchange-name"),
		amqp.WithConsumeOptionsBindingExchangeKind("fanout"),
		amqp.WithConsumeOptionsBindingExchangeDurable,
		amqp.WithConsumeOptionsQueueDurable,
	)
	return
}

func (c Consumer) StopConsuming() {
	// TODO: give this a name
	c.mqConsumer.StopConsuming("exchange-topic", true)
	c.mqConsumer.Disconnect()
}

func updateAccountStatus(d rabbitmq.Delivery) (a rabbitmq.Action) {
	// TODO: do something to update to the database
	// the client will manually refresh to get update
	return rabbitmq.Ack
}
