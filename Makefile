.PHONY: test
test:
	go test ./... -v -count=1

.PHONY: build
build:
	go build -o cmd/shortener/shortener cmd/shortener/*.go

.PHONY: run
run: build
	./cmd/shortener/shortener
