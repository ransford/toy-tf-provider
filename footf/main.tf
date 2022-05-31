terraform {
  required_providers {
    fooer = {
      version = "~> 0.0.1"
      source  = "github.com/ransford/tf-provider-fooer"
    }
  }
}

provider "fooer" {
  hostport = "localhost:8090"
}

resource "foo_thing" "mything" {
  bar = 13
}
