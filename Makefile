.PHONY: build
build:
	go build -o .bin/hamrobazar-monitor .

.PHONY: install
install:
	go install -v ./...

.PHONY: test
test:
	go test -v ./...