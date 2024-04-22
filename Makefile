.PHONY: grpc
grpc: 
	protoc --proto_path=protos protos/cheese/*.proto --go_out=. --go-grpc_out=.
