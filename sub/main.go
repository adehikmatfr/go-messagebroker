package main

import (
	"adehikmatfr/learn-go/pubsub/sub/messagebroker"
	"adehikmatfr/learn-go/pubsub/sub/messagebroker/googlepubsub"
	"adehikmatfr/learn-go/pubsub/sub/redis"
	"fmt"
)

type gb = googlepubsub.GooglePubSub
type ga = googlepubsub.GooglePubSubAdapter
type gc = googlepubsub.Config
type gaOpts = googlepubsub.AdapterOptions
type hdl = messagebroker.SubscribeMessageHandler

type testSubHdl struct{}

func (t *testSubHdl) OnProcess(msg string) {
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

	cfg := gc{
		AuthJsonPath: "assets/pubsub-credential.json",
		ProjectId:    "test-go-pub-sub",
		Strategy:     sl,
	}

	broker := &gb{
		Cfg: cfg,
	}

	rc := redis.NewRedisClient("localhost", 6379, "", 0)

	msgBroker := &ga{
		Broker: broker,
		Options: &gaOpts{
			RedisClient: rc,
		},
	}

	client := messagebroker.NewClient(msgBroker)
	thdl := &testSubHdl{}
	client.Subscribe("test-sub", thdl)
}
