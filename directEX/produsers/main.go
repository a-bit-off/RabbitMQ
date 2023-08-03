/*
PRODUCER
*/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

func main() {
	var raitingKeys = []string{"error", "info", "warning"}
	wg := sync.WaitGroup{}
	for _, key := range raitingKeys {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			loop(1, key)
		}(key)
	}

	wg.Wait()
}
func loop(interval int, routingKey string) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * time.Duration(interval))
		Producer(i, routingKey)
	}
}

func Producer(msgNumber int, routingKey string) {
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

	err = ch.Publish(
		"direct_logs",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("Message: %d  RoutingKey: %s", msgNumber, routingKey))},
	)
	if err != nil {
		log.Fatalf("Ошибка при публикации очереди: %v", err)
	}
	log.Println("Успешная публикация очереди!")

	log.Println("End producer APP")
}
