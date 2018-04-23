proto:
	@protoc --proto_path=. --go_out=. *.proto
.PHONY: proto

init:
	go get github.com/golang/dep/cmd/dep
.PHONY: init

deps:
	dep ensure
.PHONY: deps

test:
	go test -v $(go list ./...)
.PHONY: test