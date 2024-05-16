package mq

import (
	"encoding/json"
	"log"

	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
	"github.com/streadway/amqp"
)

// PublishBook publishes a new book data to the Rabbit MQ Instance
func PublishBook(queueName string, book model.Book) {
	conn, err := amqp.Dial(Server)

	if err != nil {
		log.Printf("Error connecting Rabbit MQ Instance: %s \n", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Printf("Error oppening channel: %s \n", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Error declaring queue: %s \n", err)
	}

	jsonData, err := json.Marshal(book)

	if err != nil {
		log.Printf("Error encoding JSON: %s \n", err)
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	)

	if err != nil {
		log.Printf("Error publishing message: %s \n", err)
	}

	log.Println("Successfully Published Message to Queue.")
}

// CheckMQ checks if RabbitMQ is up and running
func CheckMQ(url string) bool {
	conn, err := amqp.Dial(url)

	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
