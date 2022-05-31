terraform {
  required_providers {
    foo = {
      version = "~> 0.3.1"
      source  = "ransford.org/edu/foo"
    }
  }
}

provider "foo" {
  foo = 12
  access_key = "moo"
}

resource "foo_thing" "mine" {
  beep = "blorp"
  // fooval = 13
}

output "thing" {
  value = foo_thing.mine
}
