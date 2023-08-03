/*
PRODUCER
*/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
		Producer(i)
	}
}

func Producer(msgNumber int) {
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
			Body:        []byte(fmt.Sprintf("Message: %d", msgNumber))},
	)
	if err != nil {
		log.Fatalf("Ошибка при публикации очереди: %v", err)
	}
	log.Println("Успешная публикация очереди!")

	log.Println("End producer APP")
}
