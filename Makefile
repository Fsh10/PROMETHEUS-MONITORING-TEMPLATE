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