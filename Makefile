.PHONY: vendor 
PACKAGES = $(shell go list ./... | grep -v /vendor/)

install-tools: 
		go get -u github.com/kardianos/govendor
		go get -u github.com/alecthomas/gometalinter
		gometalinter --install --update
vendor: 
		govendor add +external	
		govendor remove +unused
build:
		go build ./...
lint: build
		gometalinter --vendor ./...
test:
		go test -v -cover $(PACKAGES)
install:
		go install
release: lint test install
