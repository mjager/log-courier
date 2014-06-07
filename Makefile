.PHONY: go-check all

export GOPATH := ${PWD}

TAGS := 'zmq zmq_4_x'

all: bin/logstash-forwarder bin/genkey

go-check:
	@go version > /dev/null || (echo "Go not found. You need to install go: http://golang.org/doc/install"; false)
	@go version | grep -q 'go version go1.[12]' || (echo "Go version 1.1.x or 1.2.x required, you have a version of go that is not supported."; false)

bin/%: src/%/*.go | go-check
	go get -d -tags $(TAGS) $*
	go install -tags $(TAGS) $*

clean:
	go clean -i ./...
