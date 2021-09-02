package ports

type QueuePort interface {
	SendMessage(topic string, message string)
	ReceiveMessage()
	Monitoring()
}
