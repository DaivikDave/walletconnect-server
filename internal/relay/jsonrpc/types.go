package jsonrpc

type Methods struct {
	Info         string
	Connect      string
	Disconnect   string
	Publish      string
	Subscribe    string
	Subscription string
	Unsubscribe  string
}

type SubscribeParams struct {
	topic string
}

type PublishParams struct {
	Topic   string
	Message string
	Ttl     float64
	Prompt  bool
}

type SubscriptionData struct {
	Topic   string
	Message string
}

type SubscriptionParams struct {
	ID   string
	Data SubscriptionData
}

type UnsubscribeParams struct {
	ID    string
	Topic string
}
