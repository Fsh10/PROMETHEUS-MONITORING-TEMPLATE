.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: generate
generate:
	@echo "Generating protobuf files..."
	@mkdir -p pkg/metrics
	@protoc \
		--go_out=pkg/metrics --go_opt=paths=import \
		--go-grpc_out=pkg/metrics --go-grpc_opt=paths=import \
		api/metrics.proto

.PHONY: build
		api/metrics.proto

.PHONY: build
build: generate
	@echo "Building application..."
	@mkdir -p pkg/metrics
	@protoc \
		--go_out=pkg/metrics --go_opt=paths=import \
		--go-grpc_out=pkg/metrics --go-grpc_opt=paths=import \
		api/metrics.proto

.PHONY: build
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/app cmd/main.go

.PHONY: build-local
build-local: generate
	@echo "Building application for local platform..."
	@mkdir -p bin
	@go build -o bin/app cmd/main.go
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/app cmd/main.go

.PHONY: build-local
build-local: generate
	@echo "Building application for local platform..."