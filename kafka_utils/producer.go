package kafka_utils

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"zph/models/util"
)

var msgProducer sarama.AsyncProducer

func init() {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForLocal

	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	//使用配置,新建一个异步生产者
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return
	}
	msgProducer = producer

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				fmt.Println("produce success")
			case err := <-p.Errors():
				fmt.Printf("err is: %+v\n", err)
			}
		}
	}(msgProducer)
}

func ProduceMsg(msg util.Message) {
	topic := "msg_topic_test"
	key := []byte(strconv.FormatUint(msg.MessageId, 10))
	content, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("ProduceMsg err ...")
	}
	m := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(content),
	}
	msgProducer.Input() <- m

}
