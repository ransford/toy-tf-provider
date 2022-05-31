A toy Terraform provider loosely based on
https://github.com/hashicorp/terraform-provider-hashicups.

To use:

    make

    go run ./cmd/server &

    cd terraform
    terraform init
    terraform apply

Now change something:

    $EDITOR main.tf   # change the foo_thing resource
    terraform plan
    terraform apply
