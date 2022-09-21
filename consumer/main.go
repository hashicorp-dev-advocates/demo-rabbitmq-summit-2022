package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	credentials, err := rabbitmqCredsGenerator()
	if err != nil {
		log.Fatal(err)
	}

	queueName := "rabbitmq-summit"
	url := fmt.Sprintf("amqp://" + credentials.Data.Username + ":" + credentials.Data.Password + "@127.0.0.1:5672/rabbitmq_summit_vhost")
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	msgs, err := channel.Consume(queueName, "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	for msg := range msgs {
		decryptedMsg, err := vaultDecryptor(msg.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("message received: " + string(decryptedMsg))
		msg.Ack(true)
	}

	defer connection.Close()

}
