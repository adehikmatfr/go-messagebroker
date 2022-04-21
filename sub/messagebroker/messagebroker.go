package messagebroker

type Messagebroker interface {
	Init() error
	Publish(m PublishMessage) error
	Subscrib(s SubscribMessage) error
}

type PublishMessage struct {
	Name    string
	Message string
	Options PublishOptions
}

type SubscribMessage struct {
	Name   string
	Listen func(msg string)
}

type PublishOptions struct {
	EnableOrdering bool
	OrderingKey    string
}

type Client struct{}

var mb Messagebroker

func setLocalMessageBroker(b Messagebroker) {
	mb = b
}

func (c *Client) Init(msgb Messagebroker) error {
	setLocalMessageBroker(msgb)
	return msgb.Init()
}

func (c *Client) Publish(m PublishMessage) error {
	return mb.Publish(m)
}

func (c *Client) Subscrib(s SubscribMessage) error {
	return mb.Subscrib(s)
}
