dev:
	go run ./cmd/web

build:
	go build ./cmd/web

run-build:
	make build
	./web -cache

test:
	go test -v -count=1 ./cmd/web
