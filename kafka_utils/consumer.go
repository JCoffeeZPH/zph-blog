package kafka_utils

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
	"zph/safe"
)

type KafkaConfig struct {
	Brokers []string
	GroupId string
	Topics  []string
}

func StartConsume(p *safe.GoPool, kafkaConfig KafkaConfig, consumerHandler *handlerConsumer, groupId string) {
	client, err := newKafkaConsumerGroupClient(kafkaConfig.Brokers, false, groupId)
	if err != nil {
		fmt.Printf("err is: %+v\n", err)
		return
	}
	consumeMsg(p, client, consumerHandler, kafkaConfig.Topics)
}

func consumeMsg(p *safe.GoPool, client sarama.ConsumerGroup, consumerHandler *handlerConsumer, topics []string) {
	if consumerHandler == nil {
		return
	}

	// consume
	p.ExecuteWithRecover(func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("consumeMsg done")
				return
			default:
				err := client.Consume(ctx, topics, consumerHandler)
				if err != nil {
					fmt.Println("consume err")
				}
				if ctx.Err() != nil {
					fmt.Println("ctx was canceled")
					return
				}
			}
		}
	}, func(err interface{}) {
		fmt.Printf("err is: %+v", err)
	})

	p.ExecuteWithRecover(func(ctx context.Context) {
		for {
			select {
			case err := <-client.Errors():
				fmt.Printf("Kafka consumer  error: %v\n", err.Error())
			case <-ctx.Done():
				return
			}

		}
	}, func(err interface{}) {
		fmt.Printf("err is: %+v", err)
	})

	p.ExecuteWithRecover(func(ctx context.Context) {
		listenSignals(ctx, client, consumerHandler)
	}, nil)
}

func newKafkaConsumerGroupClient(servers []string, fromBegin bool, groupId string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.MaxProcessingTime = 60 * time.Second
	config.Consumer.Fetch.Min = 1024
	config.Consumer.MaxWaitTime = 500 * time.Millisecond
	config.Version = sarama.V1_0_0_0
	if fromBegin {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumerGroup, err := sarama.NewConsumerGroup(servers, groupId, config)
	return consumerGroup, err
}

func listenSignals(ctx context.Context, client sarama.ConsumerGroup, consumerHandler *handlerConsumer) {
	for {
		select {
		case <-ctx.Done():
			if err := client.Close(); err != nil {
				fmt.Printf("failed to close consumer client: %v\n", err)
			}
			consumerHandler.shutdownFunc()
			return
		}
	}
}
