GOPATH:=$(CURDIR)
export GOPATH

all: go-deps build

build:
	go build -o bin/kafkaer main

install: output

get: go-deps
	go get ./...

test: go-deps
	./build_deps.sh -run_test

clean:
	go clean -i ./...
	rm -f bin/kafka-importer
	rm -rf status/

output: build
	mkdir -p output/bin
	mkdir -p output/conf
	mkdir -p output/log
	mkdir -p output/status
	cp -r conf/* output/conf/

go-deps:
	go get code.google.com/p/gcfg
	go get github.com/Shopify/sarama
	go get github.com/garyburd/redigo/redis
	go get github.com/wvanbergen/kafka/consumergroup
	go get github.com/wvanbergen/kazoo-go
