/*
CONSUMER WARNING
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

	// для перестраховки, если producer не успел создать - коммутатора
	// consumers создаст его сам
	err = ch.ExchangeDeclare(
		"direct_logs",
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка при инициализации коммутатора: %v", err)
	}
	log.Println("Успешная инициализации коммутатора!")

	// для перестраховки, если producer не успел создать очередь -
	// consumers создаст его сам
	queue, err := ch.QueueDeclare(
		"",
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка при инициализации очереди: %v", err)
	}
	log.Println("Успешная инициализации очереди!")

	if err = ch.QueueBind(
		queue.Name,
		"error",
		"direct_logs",
		false,
		nil,
	); err != nil {
		log.Fatalf("Ошибка при связке очереди (key: error): %v", err)
	}
	log.Println("Успешная связка очереди (key: error)!")

	if err = ch.QueueBind(
		queue.Name,
		"info",
		"direct_logs",
		false,
		nil,
	); err != nil {
		log.Fatalf("Ошибка при связке очереди (key: info): %v", err)
	}
	log.Println("Успешная связка очереди (key: info)!")

	if err = ch.QueueBind(
		queue.Name,
		"warning",
		"direct_logs",
		false,
		nil,
	); err != nil {
		log.Fatalf("Ошибка при связке очереди (key: warning): %v", err)
	}
	log.Println("Успешная связка очереди (key: warning)!")

	msgs, err := ch.Consume(
		queue.Name,
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
