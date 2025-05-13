BINARY_NAME = ezex-gateway
BUILD_DIR = build
CMD_DIR = ./internal/cmd/server

########################################
### Targets needed for development

gen-graphql:
	@echo "Generating graphql code..."
	@go tool gqlgen generate ./...

docker:
	docker build --tag ezex-gateway .

########################################
### Building

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

release:
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags "-s -w" -trimpath -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf $(BUILD_DIR)

########################################
### Testing

unit_test:
	@echo "Running unit tests..."
	@go test ./...

race_test:
	@echo "Running race condition tests..."
	@go test ./... -race

integration_test:
	@echo "Running integration tests..."
	@go test -tags=integration ./...

test: unit_test race_test

########################################
### Formatting the code

fmt:
	@echo "Formatting code..."
	@go tool gofumpt -l -w .

lint:
	@echo "Running lint..."
	@go tool golangci-lint  run ./... --timeout=20m0s

check: fmt lint

.PHONY: gen-graphql docker
.PHONY: build release clean
.PHONY: test unit_test race_test integration_test
.PHONY: fmt lint check
