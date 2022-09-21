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

	err = channel.ExchangeDeclare("rabbitmq_summit_exchange", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	encMsg, err := vaultEncryptor()

	message := amqp.Publishing{
		Body: []byte(encMsg.CryptoData.Ciphertext),
	}

	err = channel.Publish("rabbitmq_summit_exchange", "random-key", false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	err = channel.QueueBind(queueName, "#", "rabbitmq_summit_exchange", false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}

	defer connection.Close()

}
