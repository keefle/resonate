all: network/service.pb.go


network/service.pb.go: proto/service.proto
	protoc -I proto/ proto/service.proto --go_out=plugins=grpc:network/
