BUILD := $(shell date -u +%m%d%H%M)

build:
	@echo "Building migrator ..."
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/darwin/migrate -ldflags "-X main.BUILD=S$(BUILD)" .
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/linux/migrate -ldflags "-X main.BUILD=$(BUILD)" .

tests:
	go test .
