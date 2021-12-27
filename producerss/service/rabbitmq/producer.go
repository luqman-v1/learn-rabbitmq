package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Produce struct {
	QueueName string
}

func NewProduce(queueName string) *Produce {
	return &Produce{
		QueueName: queueName,
	}
}

func (p *Produce) QueueDeclare() {
	_, err := RabbitMQCh.QueueDeclare(
		p.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed create queue RabbitMQ ", err)
	}
}

//Queue process queue
func (p *Produce) Queue(MessageData amqp.Publishing) {
	p.QueueDeclare()
	if err := RabbitMQCh.Publish(
		"",
		p.QueueName,
		false,
		false,
		MessageData,
	); err != nil {
		log.Fatal("Failed Publish message RabbitMQ ", err)
	}
	fmt.Println("Successfully publish message")
}
