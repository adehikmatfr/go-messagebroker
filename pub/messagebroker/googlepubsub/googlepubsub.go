package googlepubsub

import (
	"adehikmatfr/learn-go/pubsub/pub/messagebroker"
	"context"
	"fmt"
	"net/url"
	"os"

	"adehikmatfr/learn-go/pubsub/pub/redis"

	"cloud.google.com/go/pubsub"
)

type GooglePubSubAdapter struct {
	Broker  *GooglePubSub
	Options *AdapterOptions
}

type AdapterOptions struct {
	RedisClient *redis.RedisClient
}

type GooglePubSub struct {
	Cfg Config
}

type Strategy struct {
	TopicName         string
	SubscriptionNames []string
}

type Config struct {
	// start from path project
	AuthJsonPath string
	ProjectId    string
	Strategy     []Strategy
}

type PublishMessage struct {
	TopicName             string
	Message               string
	EnableMessageOrdering bool
	OrderingKey           string
}

var c *pubsub.Client
var rc *redis.RedisClient

func (ga *GooglePubSubAdapter) Init() error {
	setLocalRedisClient(ga.Options.RedisClient)
	e := ga.Broker.initProccess()
	return e
}

func (ga *GooglePubSubAdapter) Publish(m messagebroker.PublishMessage) error {
	var e error
	gm := PublishMessage{
		TopicName:             m.Name,
		Message:               m.Message,
		EnableMessageOrdering: m.Options.EnableOrdering,
		OrderingKey:           m.Options.OrderingKey,
	}

	if gm.EnableMessageOrdering && ga.Options.RedisClient == nil {
		return fmt.Errorf("please config redis client first for use ordering msg")
	}

	_, e = ga.Broker.publishProccess(gm)
	if e != nil {
		return e
	}

	return nil
}

func (ga *GooglePubSubAdapter) Subscribe(name string, handler messagebroker.SubscribeMessageHandler) {
	e := ga.Broker.subscribeProccess(name, handler, ga.Options.RedisClient)
	handler.OnError(e)
}

func (g *GooglePubSub) initProccess() error {
	err := g.setEnvCredential()
	if err != nil {
		return err
	}

	ctx := context.Background()
	c, err := pubsub.NewClient(ctx, g.Cfg.ProjectId)
	//not return error if projectId not found
	if err != nil {
		return err
	}
	setLocalClient(c)

	err = g.validateStrategyList()
	if err != nil {
		return err
	}

	return nil
}

func (g *GooglePubSub) publishProccess(p PublishMessage) (string, error) {
	ctx := context.Background()
	t := c.Topic(p.TopicName)
	t.EnableMessageOrdering = p.EnableMessageOrdering
	result := t.Publish(ctx, &pubsub.Message{
		Data:        []byte(p.Message),
		OrderingKey: p.OrderingKey,
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (g *GooglePubSub) subscribeProccess(subName string, handler messagebroker.SubscribeMessageHandler, rc *redis.RedisClient) error {
	sub, e := getSubscription(subName)

	if e != nil {
		return e
	}

	ctx := context.Background()

	e = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		m.Ack()
		p := true

		if m.OrderingKey != "" && rc == nil {
			p = false
		} else if m.OrderingKey != "" && rc != nil {
			lt, _ := getLastPublish(m.OrderingKey)
			pt := int(m.PublishTime.Unix())
			if lt >= pt {
				p = false
			} else {
				setLastPublish(m.OrderingKey, pt)
			}
		}

		if p {
			str := string(m.Data)
			handler.OnProcess(str)
		}
	})

	if e != nil {
		return e
	}
	return nil
}

func (g *GooglePubSub) setEnvCredential() error {
	dir, e := os.Getwd()
	if e != nil {
		return e
	}

	dir = fmt.Sprintf("%s/%s", dir, g.Cfg.AuthJsonPath)
	url, e := url.ParseRequestURI(dir)
	if e != nil {
		return e
	}
	_, e = os.Stat(dir)
	if e != nil {
		return e
	}

	e = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", url.String())
	if e != nil {
		return e
	}

	return nil
}

func (g *GooglePubSub) validateStrategyList() error {
	for _, v := range g.Cfg.Strategy {
		_, err := getTopic(v.TopicName)
		if err != nil {
			return err
		}
		for _, sub := range v.SubscriptionNames {
			_, err = getSubscription(sub)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func setLocalClient(client *pubsub.Client) {
	c = client
}

func setLocalRedisClient(redisClient *redis.RedisClient) {
	rc = redisClient
}

func setLastPublish(key string, t int) error {
	return rc.SetInt(key, t)
}

func getLastPublish(key string) (int, error) {
	tN, e := rc.GetInt(key)
	return tN, e
}

// error topic id not found it can happen because the topic
// doesn't exist or the projectId doesn't exist yet.
func getTopic(id string) (*pubsub.Topic, error) {
	ctx := context.Background()
	t := c.Topic(id)
	ok, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not found topic id %s", id)
	}
	return t, nil
}

// error Subscription id not found it can happen because the Subscription
// doesn't exist or the projectId doesn't exist yet.
func getSubscription(id string) (*pubsub.Subscription, error) {
	ctx := context.Background()
	s := c.Subscription(id)
	ok, err := s.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not found subscription id %s", id)
	}
	return s, nil
}
