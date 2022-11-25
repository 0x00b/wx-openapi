.PHONY: all vet fmt build

ALLPROTO = $(shell find ./pb -name '*.proto' 2> /dev/null)
ALLPROTOBUF = $(ALLPROTO:%.proto=%.pb.go)
PROTOC_GIT_TAG=v1.2.0
GOLANGCI_LINT_TAG=v1.33.0
WORK_PATH=$(shell cd .. && pwd)
GENERATOR_VERSION=v1.0.0
OPENAPIYML=openapi.yml

all:
	$(MAKE) fmt
	$(MAKE) vet
	# $(MAKE) lint

protoc:
	-go get -d -u github.com/golang/protobuf/protoc-gen-go
	git -C "$(shell go env GOPATH)"/src/github.com/golang/protobuf checkout ${PROTOC_GIT_TAG}
	go install github.com/golang/protobuf/protoc-gen-go

gotools:
	go get golang.org/x/tools/cmd/goimports
	#protoc-gen-go
	-go get -d -u github.com/golang/protobuf/protoc-gen-go
	git -C "$(shell go env GOPATH)"/src/github.com/golang/protobuf checkout ${PROTOC_GIT_TAG}
	go install github.com/golang/protobuf/protoc-gen-go
	#golangci-lint
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_TAG}
	# cp $(GOPATH)/bin/golangci-lint /usr/bin/golangci-lint
#提交之前建议 make format
format:
	$(MAKE) fmt
	$(MAKE) vet

downloadlint = $(shell command -v golangci-lint>/dev/null && echo 0 || echo 1)

golangci-lint:
ifeq ($(downloadlint),1)
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
endif

lint: golangci-lint
	golangci-lint run

fmt:
	go fmt ./...

imports:
	goimports -w ./

vet:
	go vet ./...

build:
	go build ./...

test:
	go test ./...

%.pb.go:%.proto
	protoc $<  --go_out=plugins=grpc:.
proto.build: $(ALLPROTOBUF) 
proto.clean:
	-rm -f $(ALLPROTOBUF)

clean:
	go clean ./...
	-$(MAKE) proto.clean
	@echo clean finished
