terraform {
  required_providers {
    foo = {
      version = "~> 0.3.1"
      source  = "ransford.org/edu/foo"
    }
  }
}

provider "foo" {
  hostport = "localhost:8090"
  access_key = "moo"
}

resource "foo_thing" "mine" {
  bar = 31337
}

output "thing" {
  value = foo_thing.mine
}
