package rabbitmq

import (
	"adehikmatfr/learn-go/pubsub/pub/messagebroker"
	"fmt"
)

type RabbitMQAdapter struct {
	Broker *RabbitMQ
}

type RabbitMQ struct {
	Cfg Config
}

type Config struct {
	Username string
	Password string
	Server   string
}

func (ra *RabbitMQAdapter) Init() error {
	fmt.Println("comming soon")
	return nil
}

func (ra *RabbitMQAdapter) Publish(m messagebroker.PublishMessage) error {
	fmt.Println("comming soon")
	return nil
}

func (ra *RabbitMQAdapter) Subscribe(name string, handler messagebroker.SubscribeMessageHandler) {
	fmt.Println("comming soon")
}
