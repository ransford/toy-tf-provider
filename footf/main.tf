terraform {
  required_providers {
    foo = {
      version = "~> 0.3.1"
      source  = "ransford.org/edu/foo"
    }
  }
}

provider "foo" {
  access_key = "foobar"
}

resource "foo_thing" "mine" {
  foo = 5
}

output "thing" {
  value = foo_thing.mine
}
