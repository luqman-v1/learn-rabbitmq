package main

import (
	"github.com/luqman-v1/learn-rabbitmq/producer/service/rabbitmq"

	"github.com/streadway/amqp"
)

func main() {
	rabbitmq.Conn()
	p := rabbitmq.NewProduce("TestQueue")
	p.Queue(amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("hello world "),
	})

}
