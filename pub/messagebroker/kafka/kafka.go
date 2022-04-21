package rabbitmq

import (
	"adehikmatfr/learn-go/pubsub/pub/messagebroker"
	"fmt"
)

type KafkaAdapter struct {
	Broker *Kafka
}

type Kafka struct {
	Cfg Config
}

type Config struct {
}

func (ra *KafkaAdapter) Init() error {
	fmt.Println("comming soon")
	return nil
}

func (ra *KafkaAdapter) Publish(m messagebroker.PublishMessage) error {
	fmt.Println("comming soon")
	return nil
}

func (ra *KafkaAdapter) Subscribe(name string, handler messagebroker.SubscribeMessageHandler) {
	fmt.Println("comming soon")
}
