package kafka_utils

import (
	"github.com/Shopify/sarama"
)

type HandlerFunc func(msg *sarama.ConsumerMessage) error
type ShutdownFunc func() error

type handlerConsumer struct {
	handlerFunc  HandlerFunc
	shutdownFunc ShutdownFunc
}

func NewHandlerConsumer(handlerFunc HandlerFunc, shutdownFunc ShutdownFunc) *handlerConsumer {
	return &handlerConsumer{
		handlerFunc:  handlerFunc,
		shutdownFunc: shutdownFunc,
	}
}

func (consumer *handlerConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *handlerConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *handlerConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := consumer.handlerFunc(message); err == nil {
			session.MarkMessage(message, "")
		}
	}

	return nil
}
