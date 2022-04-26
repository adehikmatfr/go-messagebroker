# messagebroker package
## create new client example
1. google pubsub
```go
package main

import (
	"adehikmatfr/learn-go/pubsub/sub/messagebroker"
	"adehikmatfr/learn-go/pubsub/sub/messagebroker/googlepubsub"
	"adehikmatfr/learn-go/pubsub/sub/redis"
	"fmt"
)

type msg = messagebroker.PublishMessage
type msgOpts = messagebroker.PublishOptions
type gb = googlepubsub.GooglePubSub
type ga = googlepubsub.GooglePubSubAdapter
type gc = googlepubsub.Config
type gaOpts = googlepubsub.AdapterOptions

func main() {
	var sl []googlepubsub.Strategy
	sl = append(sl, googlepubsub.Strategy{
		TopicName:         "test",
		SubscriptionNames: []string{"test-sub"},
	})

	cfg := gc{
		AuthJsonPath: "YOUR_CREDENTIAL_PATH",
		ProjectId:    "YOUR_PROJECT_ID",
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
}
```
2. rabbitmq
```go
package main

import (
	"adehikmatfr/learn-go/pubsub/sub/messagebroker"
	"adehikmatfr/learn-go/pubsub/sub/messagebroker/googlepubsub"
	"adehikmatfr/learn-go/pubsub/sub/redis"
	"fmt"
)

type msg = messagebroker.PublishMessage
type msgOpts = messagebroker.PublishOptions
type rb = rabbitmq.RabbitMQ
type ra = rabbitmq.RabbitMQAdapter
type rc = rabbitmq.Config

func main() {
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
}
```
3. kafka
```go
package main

import (
	"adehikmatfr/learn-go/pubsub/sub/messagebroker"
	"adehikmatfr/learn-go/pubsub/sub/messagebroker/googlepubsub"
	"adehikmatfr/learn-go/pubsub/sub/redis"
	"fmt"
)

type msg = messagebroker.PublishMessage
type msgOpts = messagebroker.PublishOptions
type kb = rabbitmq.Kafka
type ka = rabbitmq.KafkaAdapter
type kc = rabbitmq.Config

func main() {
	cfg := kc{}
	broker := &kb{
		Cfg: cfg,
	}
	msgBroker := &ka{
		Broker: broker,
	}

	client := messagebroker.NewClient(msgBroker)
}
```

### publish message example
```go
    client := messagebroker.NewClient(msgBroker)
	client.Publish(msg{
		Name:    "test",
		Message: "halloword",
		Options: msgOpts{
			EnableOrdering: true,
			OrderingKey:    "test-1001:aalolxyz",
		},
	})
```

### subscribe message example
```go
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
	client := messagebroker.NewClient(msgBroker)
	thdl := &testSubHdl{}
	client.Subscribe("test-sub", thdl)
}

```

## Note
rabbitmq & kafka comming soon