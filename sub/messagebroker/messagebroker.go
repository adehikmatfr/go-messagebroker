package messagebroker

type Messagebroker interface {
	Init() error
	Publish(m PublishMessage) error
	Subscribe(name string, handler SubscribeMessageHandler)
}

type SubscribeMessageHandler interface {
	OnProccess(msg string)
	OnError(err error) error
}
type PublishMessage struct {
	Name    string
	Message string
	Options PublishOptions
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

func (c *Client) Subscribe(name string, handler SubscribeMessageHandler) {
	mb.Subscribe(name, handler)
}
