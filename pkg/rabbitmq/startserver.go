package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"response-service/pkg/config"
	"response-service/pkg/response"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ch *amqp.Channel
	q  amqp.Queue
)

func StartServer(cfg config.RabbitMQ, logic response.ResponseService) {
	fmt.Println("Starting rabbitmq")
	fmt.Println(cfg.Host + ":" + cfg.Port)
	conn, err := amqp.Dial("amqp://guest:guest@" + cfg.Host + ":" + cfg.Port)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	_, err = ch.QueueDeclare(
		"updatePostTime", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")
	createResponse, err := ch.Consume(
		"response.POST", // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	deleteResponses, err := ch.Consume(
		"response.DELETE", // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range createResponse {
			var response response.CreateResponse
			if err := json.NewDecoder(bytes.NewReader(d.Body)).Decode(&response); err != nil {
				fmt.Println("Unmarshal went wrong")
				return
			}
			response.Uuid = uuid.NewString()
			logic.CreateResponse(context.Background(), response)
			UpdatePostTime(response.PostId)
		}
	}()
	go func() {
		for d := range deleteResponses {
			logic.DeleteResponses(context.Background(), string(d.Body))
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

func UpdatePostTime(postId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	bytes := []byte(postId)
	fmt.Printf("bytes: %v\n", bytes)
	err := ch.PublishWithContext(ctx,
		"",               // exchange
		"updatePostTime", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s\n", postId)
}
