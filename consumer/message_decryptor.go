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

type Plaintext struct {
	CryptoData CryptoData `json:"data"`
}

type CryptoData struct {
	Plaintext string `json:"plaintext"`
}

func vaultDecryptor(Ct []byte) (plaintext []byte, err error) {

	vaultEndpoint := os.Getenv("VAULT_ADDR")
	if vaultEndpoint == "" {
		log.Fatal("VAULT_ADDR environment variable not set.\n")
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("VAULT_TOKEN environment variable not set.\n")
	}

	payload := `{"ciphertext": "` + string(Ct) + `"}`

	url := vaultEndpoint + "/v1/transit/decrypt/rabbitmq_summit"
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

	var Pt Plaintext
	err = json.Unmarshal(body, &Pt)

	plaintext, err = b64.StdEncoding.DecodeString(Pt.CryptoData.Plaintext)
	if err != nil {
		log.Fatal(err)
	}

	return plaintext, err

}
