# Basic go commands
GOCMD=CGO_ENABLED=0 go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=build/resonate
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_FREEBSD=$(BINARY_NAME)_freebsd
BINARY_ARM=$(BINARY_NAME)_linux_arm

all: test build-local

build-local:
	$(GOCLEAN)
	$(GOBUILD) -o $(BINARY_NAME) -v

generate:
	protoc -I proto/ proto/*.proto --go_out=plugins=grpc:network/

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)
	rm -f $(BINARY_FREEBSD)
	rm -f $(BINARY_ARM)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop


release: build-linux build-freebsd build-arm

# Cross compilation
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v

build-freebsd:
	GOOS=freebsd GOARCH=amd64 $(GOBUILD) -o $(BINARY_FREEBSD) -v

build-arm:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BINARY_ARM) -v

docker-test:
	podman build -t resonate-test -f Dockerfile-test
	podman run -it --cap-add SYS_ADMIN -v  /lib/modules:/lib/modules --device /dev/fuse resonate-test

