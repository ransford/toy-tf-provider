terraform {
  required_providers {
    foo = {
      version = "~> 0.0.1"
      source  = "ransford.org/edu/foo"
    }
  }
}

provider "foo" {
  hostport = "localhost:8090"
}

resource "foo_thing" "mything" {
  bar = 13
}
