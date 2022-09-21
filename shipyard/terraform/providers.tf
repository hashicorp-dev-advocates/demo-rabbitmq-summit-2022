terraform {
  required_providers {

    rabbitmq = {
      source = "cyrilgdn/rabbitmq"
      version = "1.7.0"
    }

    vault = {
      source = "hashicorp/vault"
      version = "3.8.2"
    }

  }
}

provider "rabbitmq" {
    insecure = true
}

provider "vault" {
  # Configuration options
}