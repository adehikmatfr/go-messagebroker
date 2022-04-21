package main

import (
	"adehikmatfr/learn-go/pubsub/sub/messagebroker"
	"adehikmatfr/learn-go/pubsub/sub/messagebroker/googlepubsub"
	"fmt"
)

type client = messagebroker.Client
type broker = googlepubsub.GooglePubSub
type adapter = googlepubsub.GooglePubSubAdapter
type config = googlepubsub.Config
type sub = messagebroker.SubscribMessage

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
	client.Subscrib(sub{
		Name:   "test-sub",
		Listen: func(msg string) { fmt.Println(msg) },
	})
}
