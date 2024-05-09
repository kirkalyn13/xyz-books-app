package mq

const (
	// Server is the URL for Rabbit MQ instance.
	// Note that this is only for demo purposes and should not be included in the code.
	Server = "amqp://guest:guest@localhost:5672/"

	// BookQueue is the queue name for the books API
	BookQueue = "xyz-books"
)
