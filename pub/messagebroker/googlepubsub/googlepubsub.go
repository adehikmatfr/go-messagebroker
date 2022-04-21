package googlepubsub

import (
	"adehikmatfr/learn-go/pubsub/pub/messagebroker"
	"context"
	"fmt"
	"net/url"
	"os"

	"cloud.google.com/go/pubsub"
)

type GooglePubSubAdapter struct {
	GooglePubsubBroker *GooglePubSub
}

type GooglePubSub struct {
	Cfg Config
}

type Strategy struct {
	TopicName         string
	SubscriptionNames []string
}

type Config struct {
	AuthJsonPath string
	ProjectId    string
	Strategy     []Strategy
}

type PublishMessage struct {
	TopicName             string
	Message               string
	EnableMessageOrdering bool
}

var c *pubsub.Client

func (ga *GooglePubSubAdapter) Init() error {
	e := ga.GooglePubsubBroker.initProccess()
	return e
}

func (ga *GooglePubSubAdapter) Publish(m messagebroker.PublishMessage) error {
	gm := PublishMessage{
		TopicName:             m.Name,
		Message:               m.Message,
		EnableMessageOrdering: m.Options.EnableOrdering,
	}
	_, e := ga.GooglePubsubBroker.publishProccess(gm)
	return e
}

func (ga *GooglePubSubAdapter) Subscribe(name string, handler messagebroker.SubscribeMessageHandler) {
	e := ga.GooglePubsubBroker.subscribeProccess(name, handler)
	handler.OnError(e)
}

func (g *GooglePubSub) initProccess() error {
	err := g.setCredentialEnv()
	if err != nil {
		return err
	}

	ctx := context.Background()
	c, err := pubsub.NewClient(ctx, g.Cfg.ProjectId)
	//not return error if projectId not found
	if err != nil {
		return err
	}
	setLocalCleint(c)

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
		Data: []byte(p.Message),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (g *GooglePubSub) subscribeProccess(subName string, handler messagebroker.SubscribeMessageHandler) error {
	sub, e := getSubscription(subName)

	if e != nil {
		return e
	}

	ctx := context.Background()

	e = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		// TODO: Handle message.
		// NOTE: May be called concurrently; synchronize access to shared memory.
		m.Ack()
		str := string(m.Data)
		handler.OnProcess(str)
	})

	if e != nil {
		return e
	}
	return nil
}

func (g *GooglePubSub) setCredentialEnv() error {
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

func setLocalCleint(client *pubsub.Client) {
	c = client
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
