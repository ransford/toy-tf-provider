MODULE   := ransford.org/edu/foo
MODVER   := 0.3.1
ARCH     := $(shell uname -m)
OS       := $(shell uname -s | tr A-Z a-z)
PROVPATH := ~/.terraform.d/plugins/$(MODULE)/$(MODVER)/$(OS)_$(ARCH)/terraform-provider-foo

all: clean build init

build:
	go build -o $(PROVPATH)

init:
	cd footf && terraform init

clean:
	cd footf && rm -rf .terraform .terraform.lock.hcl
	$(RM) server
