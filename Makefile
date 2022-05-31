MODULE   := $(shell grep \^module go.mod | cut -d' ' -f2)
MODVER   := 0.0.1
ARCH     := $(shell uname -m)
OS       := $(shell uname -s | tr A-Z a-z)
PROVPATH := ~/.terraform.d/plugins/$(MODULE)/$(MODVER)/$(OS)_$(ARCH)/terraform-provider-foo

.PHONY: build
build:
	go build -o $(PROVPATH)
