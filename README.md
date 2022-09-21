# RabbitMQ Summit London 2022 - Zero Trust RabbitMQ

As demonstrated by [Rob Barnes (DevOps Rob)](https://github.com/devops-rob) during the Zero Trust RabbitMQ Talk at the [RabbitMQ Summit London](https://rabbitmqsummit.com/) on 16th September 2022.

## Demo overview

This demo written in [Go](https://go.dev/), uses a Zero Trust Security approach to securing application access to RabbitMQ as well as protecting our application data by encrypting it at rest.

The application is made of two services:

1. Message consumer - This is a listening service which obtains RabbitMQ credentials from Vault with least priviledged access to consume messages encrypted by Vault's Transit engine from a queue named `rabbitmq-summit`. This service then decrypts the messages using the cryptographic key.
2. Message producer - This service obtains RabbitMQ credentials from Vault with least priviledged access to publish messages encrypted by Vault's Transit engine to a queue named `rabbitmq-summit`.

## Running the demo

To configure Vault and RabbitMQ, you will need to run this [Terraform code](/shipyard/terraform/) which will provision the follwing resources:
1. RabbitMQ vhost
2. RabbitMQ exchange
3. Message queue
4. RabbitMQ user for Vault to broker identity on behalf of RabbitMQ
5. Secret engine in Vault for RabbitMQ
6. Vault role for the Message Producer service
7. Vault role for the Message Consumer service
8. Transit secrets engine
9. Transit cryptographic key
10. Vault policy for Message Consumer
11. Vault policy for Message Producer

### Generating Vault credentials for the services

Each service will require a Vault token with their respective Vault policies.

The following command will generate a token for the Message Consumer service:

```sh
vault token create -policy=rabbitmq-consumer
```

Copy the token from this command output and run the following commands to set the environment variables in the terminal session that will run the Message Consumer service"

```sh
export VAULT_ADDR="http://vault.container.shipyard.run:8200"
export VAULT_TOKEN="<insert copied token here>"
```

The same will need to be done for the Message Producer service:

```sh
vault token create -policy=rabbitmq-producer
```

Copy the token from this command output and run the following commands to set the environment variables in the terminal session that will run the Message Producer service"

```sh
export VAULT_ADDR="http://vault.container.shipyard.run:8200"
export VAULT_TOKEN="<insert copied token here>"
```

### Starting the services

Change directory in each of the terminals to the respective service and run them.

For the Message Consumer service:

```sh
cd consumer
go run .
```

For the Message Producer service:

```sh
cd consumer
go run . <insert message here>
```

When you run the Producer service, the message you entered will be displayed in the Consumer service terminal
