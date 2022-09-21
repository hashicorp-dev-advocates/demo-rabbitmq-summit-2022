package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Ciphertext struct {
	CryptoData CryptoData `json:"data"`
}

type CryptoData struct {
	Ciphertext string `json:"ciphertext"`
}

func vaultEncryptor() (CT Ciphertext, err error) {

	argsWithoutProg := os.Args[1:]

	msg := b64.StdEncoding.EncodeToString([]byte(argsWithoutProg[0]))

	vaultEndpoint := os.Getenv("VAULT_ADDR")
	if vaultEndpoint == "" {
		log.Fatal("VAULT_ADDR environment variable not set.\n")
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("VAULT_TOKEN environment variable not set.\n")
	}

	payload := `{"plaintext": "` + msg + `"}`

	url := vaultEndpoint + "/v1/transit/encrypt/rabbitmq_summit"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
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

	err = json.Unmarshal(body, &CT)

	return CT, err

}
