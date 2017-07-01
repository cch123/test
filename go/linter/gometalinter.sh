#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export GOBIN=$(GOPATH)/bin
export CGO_ENABLED=0
export GOOS=linux

all: lint

lint:
    which gometalinter || (go get -u -v github.com/alecthomas/gometalinter && gometalinter --install)
    gometalinter --vendor --skip=vendor/ --exclude=vendor \
    --disable-all \
    --enable=gofmt \
    --enable=vet --enable=vetshadow \
    --enable=gocyclo --cyclo-over=24 \
    --enable=golint \
    --enable=ineffassign \
    --enable=misspell \
    --deadline=5m \
    --concurrency=1 \
    ./...

format:
    which goimports || go get -u -v golang.org/x/tools/cmd/goimports
    find $(ROOT)/ -type f -name "*.go" | grep -v $(ROOT)/vendor | xargs --max-args=1 --replace=R goimports -w R
    find $(ROOT)/ -type f -name "*.go" | grep -v $(ROOT)/vendor | xargs --max-args=1 --replace=R gofmt -s -w R
