package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Queue struct {
	QueueName, Consumer                 string
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                amqp.Table
}

func NewQueue(q Queue) *Queue {
	return &Queue{
		QueueName: q.QueueName,
		Consumer:  q.Consumer,
		AutoAck:   q.Exclusive,
		Exclusive: q.Exclusive,
		NoLocal:   q.NoLocal,
		NoWait:    q.NoWait,
		Args:      q.Args,
	}
}

func (q *Queue) Run() {
	messages, err := RabbitMQCh.Consume(
		q.QueueName,
		q.Consumer,
		q.AutoAck,
		q.Exclusive,
		q.NoLocal,
		q.NoLocal,
		q.Args,
	)
	if err != nil {
		log.Fatal("Failed Consume channel RabbitMQ ", err)
	}
	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("Recived message : %s\n", message.Body)
			if err := message.Ack(false); err != nil {
				log.Fatal("Failed Ack message RabbitMQ ", err)
			}
		}
	}()
	fmt.Println("Successfully connected to our rabbitMq ...")
	fmt.Println("[*] - Waiting for messages")
	<-forever
}
