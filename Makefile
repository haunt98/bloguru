.PHONY: help gen

help:
	@echo "Reading code is the way"

gen:
	protoc -I . \
	--go_out . --go_opt module=github.com/haunt98/bloguru \
	--go-grpc_out . --go-grpc_opt module=github.com/haunt98/bloguru \
	--grpc-gateway_out . \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt module=github.com/haunt98/bloguru \
	--grpc-gateway_opt grpc_api_configuration=proto/author/v1/config.yaml \
	proto/author/v1/author.proto