fmt:
	go fmt ./...

build:
	go build -o bin/offvsix cmd/offvsix/main.go

release: fmt build
	strip bin/offvsix

run:
	go run cmd/offvsix/main.go -- golang.go