
proto:
	protoc -I . --go_out ./pb/generated --go-grpc_out ./pb/generated pb/model/*.proto
