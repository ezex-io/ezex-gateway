BINARY_NAME = ezex-gateway
BUILD_DIR = build
CMD_DIR = ./internal/cmd/server

########################################
### Targets needed for development

graphql:
	@echo "Generating graphql code..."
	@go tool gqlgen generate ./...

docker:
	docker build --tag ezex-gateway .

mock:
	mockgen -source=./internal/port/authenticator.go 	-destination=./internal/mock/mock_authenticator.go	-package=mock
	mockgen -source=./internal/port/cache.go 			-destination=./internal/mock/mock_cache.go 			-package=mock
	mockgen -source=./internal/port/notification.go 	-destination=./internal/mock/mock_notification.go 	-package=mock
	mockgen -source=./internal/port/users.go 			-destination=./internal/mock/mock_users.go 			-package=mock

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

test: unit_test

########################################
### Formatting the code

fmt:
	@echo "Formatting code..."
	@go tool gofumpt -l -w .

lint:
	@echo "Running lint..."
	@go tool golangci-lint  run ./... --timeout=20m0s

check: fmt lint

.PHONY: graphql docker mock
.PHONY: build release clean
.PHONY: test unit_test race_test integration_test
.PHONY: fmt lint check
