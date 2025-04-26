CGO_ENABLED = 0 # statically linked = 0
TARGETOS=linux
ifeq ($(OS),Windows_NT)
    TARGETOS := Windows
else
    TARGETOS := $(shell sh -c 'uname 2>/dev/null || echo Unknown' | tr '[:upper:]' '[:lower:]')
endif
TARGETARCH = amd64

.DEFAULT_GOAL = help
PID = /tmp/serving.pid
DEVELOPER_UID     ?= $(shell id -u)
DC = docker compose

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

.PHONY: build
build: ## Build dev image
	@$(DC) build

.PHONY: go-init
go-init: up
	@$(DC) exec godev sh -c "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"

.PHONY: php-init
php-init:
	@$(DC) exec phpdev sh -c "composer install"
	@$(DC) exec phpdev sh -c "curl -L https://github.com/roadrunner-server/roadrunner/releases/download/v2024.3.5/protoc-gen-php-grpc-2024.3.5-linux-amd64.tar.gz | tar -xzO protoc-gen-php-grpc-2024.3.5-linux-amd64/protoc-gen-php-grpc > /app/bin/protoc-gen-php-grpc && chmod +x /app/bin/protoc-gen-php-grpc"

.PHONY: up
up: ## Start application dev containers
	@$(DC) up -d

.PHONY: down
down: ## Remove application dev containers
	@$(DC) down

.PHONY: php-shell
php-shell: ## Enter PHP application dev container
	@$(DC) exec -it phpdev bash

.PHONY: php-protoc
php-protoc: ## Create Hermes gRPC classes
	@$(DC) exec -it phpdev sh -c "protoc --php_out=src/Infrastructure/Grpc --grpc_out=src/Infrastructure/Grpc --plugin=protoc-gen-grpc=bin/protoc-gen-php-grpc --proto_path=proto proto/hermes.proto"
	@$(DC) exec -it phpdev sh -c "mv -f src/Infrastructure/Grpc/App/Infrastructure/Grpc/Hermes/*.php src/Infrastructure/Grpc/Hermes/ && rm -fr src/Infrastructure/Grpc/App"
	@$(DC) exec -it phpdev sh -c "find src/Infrastructure/Grpc -type f -name '*.php' -exec sed -i 's|^namespace GPBMetadata;|namespace App\\\\Infrastructure\\\\Grpc\\\\GPBMetadata;|g' {} +"
	@echo "You have fix HermesRequest and HermesResponse (__construct)!!!"

.PHONY: go-protoc
go-protoc:
	@$(DC) exec -it godev sh -c 'protoc --proto_path=./proto --go_out=. --go_opt=paths=source_relative,Mhermes.proto=hermes/gen/hermes --go-grpc_out=. --go-grpc_opt=paths=source_relative,Mhermes.proto=hermes/gen/hermes ./proto/hermes.proto'
	@$(DC) exec -it godev sh -c 'mv hermes_grpc.pb.go gen/hermes && mv hermes.pb.go gen/hermes'

#.PHONY: php-serve
#php-serve:
#	@$(DC) exec -it phpdev sh -c 'bin/rr serve -c .rr.yaml'

.PHONY: go-shell
go-shell: ## Enter GO application dev container
	@$(DC) exec -it godev bash

.PHONY: go-build
go-build: ## Build dev application (go build)
	@$(DC) exec godev sh -c "env CGO_ENABLED=${CGO_ENABLED} GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o bin/hermes ./"

#.PHONY: go-init
#go-init: up
#	@$(DC) exec godev sh -c "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
