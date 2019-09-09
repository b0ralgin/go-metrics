NAME := metrics
VERSION := $(shell git rev-parse --abbrev-ref HEAD)

LD_FLAGS += -X main.name=$(NAME)
LD_FLAGS += -X main.version=$(VERSION)

.PHONY:all
all: clean modtidy modverify moddownload lint golangcilint test build

.PHONY: build
build:
	go build -ldflags '$(LD_FLAGS)' -o bin/$(NAME) metrics/cmd/.

.PHONY: clean
clean:
	rm -f bin/$(NAME)

.PHONY:modtidy
modtidy:
	go mod tidy

.PHONY: modverify
modverify:
	go mod verify

.PHONY:moddownload
moddownload:
	go mod download

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: golangcilint
golangcilint:
	golangci-lint run --enable-all --deadline=15m

.PHONY: test
test:
	go test -v -cover ./...
