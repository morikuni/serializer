proto:
	@protoc --proto_path=. --go_out=. *.proto
.PHONY: proto