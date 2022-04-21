package main

import (
	"adehikmatfr/learn-go/pubsub/pub/messagebroker"
	"adehikmatfr/learn-go/pubsub/pub/messagebroker/googlepubsub"
	"adehikmatfr/learn-go/pubsub/pub/messagebroker/rabbitmq"
	"fmt"
)

type msg = messagebroker.PublishMessage
type msgOpts = messagebroker.PublishOptions
type gb = googlepubsub.GooglePubSub
type ga = googlepubsub.GooglePubSubAdapter
type gc = googlepubsub.Config
type gaOpts = googlepubsub.AdapterOptions
type rb = rabbitmq.RabbitMQ
type ra = rabbitmq.RabbitMQAdapter
type rc = rabbitmq.Config

func main() {
	// var sl []googlepubsub.Strategy
	// sl = append(sl, googlepubsub.Strategy{
	// 	TopicName:         "test",
	// 	SubscriptionNames: []string{"test-sub"},
	// })

	// cfg := gc{
	// 	AuthJsonPath: "assets/pubsub-credential.json",
	// 	ProjectId:    "test-go-pub-sub",
	// 	Strategy:     sl,
	// }

	// broker := &gb{
	// 	Cfg: cfg,
	// }

	// rc := redis.NewRedisClient("localhost", 6379, "", 0)

	cfg := rc{
		Username: "guest",
		Password: "guest",
		Server:   "localhost:5672",
	}
	broker := &rb{
		Cfg: cfg,
	}
	msgBroker := &ra{
		Broker: broker,
	}

	client := messagebroker.NewClient(msgBroker)
	for i := 0; i < 10; i++ {
		client.Publish(msg{
			Name:    "test",
			Message: fmt.Sprintf("88888-%d", i),
			Options: msgOpts{
				EnableOrdering: true,
				OrderingKey:    "test-1001:aalolxyz",
			},
		})
	}

}
