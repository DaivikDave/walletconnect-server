package relay

type LegacySocketMessage struct {
	Topic   string
	Type    string
	Payload string
	Silent  bool
}

type Notification struct {
	Topic   string
	Webhook string
}
