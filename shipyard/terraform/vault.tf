# resource "vault_mount" "rabbitmq" {
#   path = "rabbitmq"
#   type = "rabbitmq"
# }

resource "vault_rabbitmq_secret_backend" "rabbitmq" {
  connection_uri = "http://rabbitmq.container.shipyard.run:15672"
  username       = rabbitmq_user.vault.name
  password       = rabbitmq_user.vault.password
  #   path           = vault_mount.rabbitmq.path
}

resource "vault_rabbitmq_secret_backend_role" "producer" {
  backend = vault_rabbitmq_secret_backend.rabbitmq.path
  name    = "producer"

  tags = "management"

  vhost {
    host      = rabbitmq_vhost.rabbitmq_summit_vhost.name
    configure = ".*"
    read      = ".*"
    write     = ".*"
  }

  vhost_topic {
    vhost {
      topic = rabbitmq_exchange.rabbitmq_summit_exchange.name
      read  = ".*"
      write = ".*"
    }

    host = rabbitmq_vhost.rabbitmq_summit_vhost.name
  }
}

resource "vault_rabbitmq_secret_backend_role" "consumer" {
  backend = vault_rabbitmq_secret_backend.rabbitmq.path
  name    = "consumer"

  tags = "management"

  vhost {
    host      = rabbitmq_vhost.rabbitmq_summit_vhost.name
    configure = ".*"
    read      = ".*"
    write     = ""
  }

  vhost_topic {
    vhost {
      topic = rabbitmq_exchange.rabbitmq_summit_exchange.name
      read  = ".*"
      write = ""
    }

    host = rabbitmq_vhost.rabbitmq_summit_vhost.name
  }
}

resource "vault_policy" "producer" {
  name = "rabbitmq-producer"

  policy = <<EOP
path "rabbitmq/creds/producer" {
  capabilities = ["read"]
}

path "transit/encrypt/rabbitmq_summit" {
  capabilities = ["update"]
}
EOP
}

resource "vault_policy" "consumer" {
  name = "rabbitmq-consumer"

  policy = <<EOP
path "rabbitmq/creds/consumer" {
  capabilities = ["read"]
}

path "transit/decrypt/rabbitmq_summit" {
  capabilities = ["update"]
}

EOP
}
resource "vault_mount" "transit" {
  path = "transit"
  type = "transit"
}

resource "vault_transit_secret_backend_key" "rabbitmq_summit" {
  backend          = vault_mount.transit.path
  name             = "rabbitmq_summit"
  deletion_allowed = true
}