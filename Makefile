all:
	protoc -I proto/ proto/service.proto --go_out=plugins=grpc:network/
