# supplier-hub

Hosts suppliers in a singular monolith.

## Setup

### Environment setup

    cp configs/.env-default configs/.env

### Install packages

    go mod download

### Start service

    go run cmd/main.go

### Generate resources

    ./bin/gen-resources.sh
