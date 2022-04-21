package main

import (
	"adehikmatfr/learn-go/pubsub/pub/messagebroker"
	"adehikmatfr/learn-go/pubsub/pub/messagebroker/googlepubsub"
)

type client = messagebroker.Client
type broker = googlepubsub.GooglePubSub
type adapter = googlepubsub.GooglePubSubAdapter
type config = googlepubsub.Config
type msg = messagebroker.PublishMessage
type msgOpts = messagebroker.PublishOptions

func main() {
	var sl []googlepubsub.Strategy
	sl = append(sl, googlepubsub.Strategy{
		TopicName:         "test",
		SubscriptionNames: []string{"test-sub"},
	})
	cfg := config{
		AuthJsonPath: "assets/pubsub-credential.json",
		ProjectId:    "test-go-pub-sub",
		Strategy:     sl,
	}
	client := &client{}
	googlePubSubBroker := &broker{
		Cfg: cfg,
	}
	googlePubSubAdapter := &adapter{
		GooglePubsubBroker: googlePubSubBroker,
	}
	client.Init(googlePubSubAdapter)
	client.Publish(msg{
		Name:    "test",
		Message: "hallo word",
		Options: msgOpts{EnableOrdering: true},
	})
}
