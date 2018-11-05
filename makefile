migrate:
	migrate -database "mysql://test:test@tcp(127.0.0.1:8889)/weatherdb" -path ./migrations up

start: build
		./main

build:
	go build -v ./cmd/main.go

lint-fix:
	go vet ./...
	go fmt ./...

deps:
	go get

test:
	go test ./...

clean:
	rm ./main