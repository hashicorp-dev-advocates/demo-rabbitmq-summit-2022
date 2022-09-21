resource "rabbitmq_user" "vault" {
  name     = "vault"
  password = "password"
  tags     = ["administrator", "management"]
}

resource "rabbitmq_permissions" "vault" {
  user  = rabbitmq_user.vault.name
  vhost = rabbitmq_vhost.rabbitmq_summit_vhost.name

  permissions {
    configure = ".*"
    write     = ".*"
    read      = ".*"
  }
}

resource "rabbitmq_vhost" "rabbitmq_summit_vhost" {
  name = "rabbitmq_summit_vhost"
}

resource "rabbitmq_exchange" "rabbitmq_summit_exchange" {
  name  = "rabbitmq_summit_exchange"
  vhost = rabbitmq_vhost.rabbitmq_summit_vhost.name

  settings {
    type        = "topic"
    durable     = true
    auto_delete = false
  }
}

resource "rabbitmq_queue" "rabbitmq-summit" {
  name = "rabbitmq-summit"
  vhost = rabbitmq_vhost.rabbitmq_summit_vhost.name
  
  settings {
    durable = true
    auto_delete = false
  }
}