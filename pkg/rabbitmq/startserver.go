package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"response-service/pkg/config"
	"response-service/pkg/response"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartServer(cfg config.RabbitMQ, logic response.ResponseService) {
	fmt.Println("Starting rabbitmq")
	fmt.Println(cfg.Host + ":" + cfg.Port)
	conn, err := amqp.Dial("amqp://guest:guest@" + cfg.Host + ":" + cfg.Port)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			fmt.Printf("d.UserId: %v\n", d.UserId)
			var response response.CreateResponse
			if err := json.NewDecoder(bytes.NewReader(d.Body)).Decode(&response); err != nil {
				fmt.Println("Unmarshal went wrong")
				return
			}
			response.Uuid = uuid.NewString()
			logic.CreateResponse(context.Background(), response)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
