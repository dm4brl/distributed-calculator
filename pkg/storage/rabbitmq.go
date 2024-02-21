package rabbitmq

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/yourusername/yourproject/config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(cfg *config.Config) (*RabbitMQ, error) {
	conn, err := amqp.Dial(cfg.Server.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	queue, err := channel.QueueDeclare(
		"calculator", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err)
	}

	return &RabbitMQ{conn: conn, channel: channel, queue: queue}, nil
}

func (r *RabbitMQ) Calculate(expression string) (float64, error) {
	err := r.channel.Publish(
		"",           // exchange
		r.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(expression),
		})
	if err != nil {
		return 0, fmt.Errorf("failed to publish a message: %v", err)
	}

	msgs, err := r.channel.Consume(
		r.queue.Name, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return 0, fmt.Errorf("failed to register a consumer: %v", err)
	}

	for d := range msgs {
		result, err := strconv.ParseFloat(string(d.Body), 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse result: %v", err)
		}

		return result, nil
	}

	return 0, errors.New("no result received")
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}
