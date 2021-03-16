IGNORED_FOLDER=.ignore
MODULE_NAME := $(shell go list -m)

all: install tools lint build

.PHONY: install
install:
	@go mod download

.PHONY: build
build: test
	@go build -a -ldflags '-extldflags "-static"' ./...

test:
	@go test -count=1 ./...

.PHONY: lint
lint:
	golint ./...

.PHONY: tools
tools:
	go get -u golang.org/x/lint/golint

.PHONY: clean
clean:
	@rm -rf ${IGNORED_FOLDER} 

.PHONY: fclean
fclean: clean
	@rm -rf ${BIN_FOLDER}
