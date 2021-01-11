.PHONY: help mod gen clean lint

help:
	@echo "Reading code is the way"

mod:
	go mod tidy
	go install \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        github.com/bufbuild/buf/cmd/buf \
        github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
        github.com/bufbuild/buf/cmd/protoc-gen-buf-lint

gen:
	protoc -I . \
	--go_out . --go_opt module=github.com/haunt98/bloguru \
	--go-grpc_out . --go-grpc_opt module=github.com/haunt98/bloguru \
	--grpc-gateway_out . \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt module=github.com/haunt98/bloguru \
	--grpc-gateway_opt grpc_api_configuration=proto/author/v1/config.yaml \
	proto/author/v1/author.proto

clean:
	rm -rf gen/

lint:
	buf lint
	golangci-lint run ./...
