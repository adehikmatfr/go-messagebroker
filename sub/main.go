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
type hdl = messagebroker.SubscribeMessageHandler

type testSubHdl struct{}

func (t *testSubHdl) OnProccess(msg string) {
	fmt.Println(msg)
}

func (t *testSubHdl) OnError(err error) error {
	fmt.Println(err)
	return nil
}

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
	thdl := &testSubHdl{}
	client.Subscribe("test-sub", thdl)
}
