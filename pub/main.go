package main

import (
	"adehikmatfr/learn-go/pubsub/pub/messagebroker"
	"adehikmatfr/learn-go/pubsub/pub/messagebroker/googlepubsub"
)

type client = messagebroker.Client
type msg = messagebroker.PublishMessage
type msgOpts = messagebroker.PublishOptions
type gb = googlepubsub.GooglePubSub
type ga = googlepubsub.GooglePubSubAdapter
type gc = googlepubsub.Config

// type rb = rabbitmq.RabbitMQ
// type ra = rabbitmq.RabbitMQAdapter
// type rc = rabbitmq.Config

func main() {
	var sl []googlepubsub.Strategy
	sl = append(sl, googlepubsub.Strategy{
		TopicName:         "test",
		SubscriptionNames: []string{"test-sub"},
	})
	cfg := gc{
		AuthJsonPath: "assets/pubsub-credential.json",
		ProjectId:    "test-go-pub-sub",
		Strategy:     sl,
	}
	broker := &gb{
		Cfg: cfg,
	}
	msgBroker := &ga{
		Broker: broker,
	}

	// cfg := rc{
	// 	Username: "guest",
	// 	Password: "guest",
	// 	Server:   "localhsot:5672",
	// }
	// broker := &rb{
	// 	Cfg: cfg,
	// }
	// msgBroker := &ra{
	// 	Broker: broker,
	// }

	client := &client{}
	client.Init(msgBroker)
	client.Publish(msg{
		Name:    "test",
		Message: "hallo word",
		Options: msgOpts{EnableOrdering: true},
	})
}
