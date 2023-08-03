/*
CONSUMER
*/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	log.Println("Start consumers APP")

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
	log.Println("Успешное открытие канала!")
	defer ch.Close()

	// для перестраховки, если producer не успел создать очередь -
	// consumers создаст его сам
	_, err = ch.QueueDeclare(
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

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка при считывании с канала: %v", err)
	}
	log.Println("Успешная считывание с канала!")

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			fmt.Println(string(m.Body))
		}
	}()

	log.Println("Успешное подключение к интерфейсу RabbitMQ")
	log.Println("Ожидается ввод")
	<-forever

	log.Println("End consumers APP")

}
