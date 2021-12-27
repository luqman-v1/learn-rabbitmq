package main

import "github.com/luqman-v1/learn-rabbitmq/consumer/service/rabbitmq"

func main() {
	rabbitmq.Conn()
	queue := rabbitmq.NewQueue(rabbitmq.Queue{
		QueueName: "TestQueue",
		Consumer:  "",
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	})
	queue.Run()
}
