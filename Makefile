.PHONY: plugin test

plugin:
	@go build -buildmode=plugin plugin/gobreakselectinfor.go

test:
	@go test -v ./...