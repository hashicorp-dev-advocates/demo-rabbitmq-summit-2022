package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RabbitCreds struct {
	Data Data `json:"data"`
}

type Data struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func rabbitmqCredsGenerator() (RabbitmqCreds RabbitCreds, err error) {
	vaultEndpoint := os.Getenv("VAULT_ADDR")
	if vaultEndpoint == "" {
		log.Fatal("VAULT_ADDR environment variable not set.\n")
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("VAULT_TOKEN environment variable not set.\n")
	}

	url := vaultEndpoint + "/v1/rabbitmq/creds/consumer"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("X-Vault-Token", vaultToken)
	request.Header.Add("accept", "*/*")

	resp, err := Client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	err = json.Unmarshal(body, &RabbitmqCreds)

	return RabbitmqCreds, err
}
