package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const AMQP_URL = "amqp://guest:guest@127.0.0.1:5672/"

var RabbitMQCh *amqp.Channel

func Conn() {
	fmt.Println("RabbitMQ Starting ...")
	conn, err := amqp.Dial(AMQP_URL)
	if err != nil {
		log.Fatal("Failed Connection to RabbitMQ ", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed open channel RabbitMQ ", err)
	}
	RabbitMQCh = ch
}
