/*
PRODUCER
*/
package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	log.Println("Start producer APP")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Ошибка при подключении к rabbitMQ: %v", err)
	}
	log.Println("Успешное подключение к rabbitMQ!")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка при открытии канала: %v", err)
	}

	queue, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка при инициализации очереди: %v", err)
	}
	log.Println("Успешная инициализации очереди!")
	log.Println("Очередь:", queue)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello RabbitMQ!")},
	)
	if err != nil {
		log.Fatalf("Ошибка при публикации очереди: %v", err)
	}
	log.Println("Успешная публикация очереди!")

	log.Println("End producer APP")
}
