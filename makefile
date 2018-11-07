migrate:
	migrate -database "${DATABASE_URL}" -path ./migrations up

start: build
		./main

build:
	go build -v ./cmd/main.go

lint-fix:
	go vet ./...
	go fmt ./...

deps:
	dep ensure

test:
	go test ./...

clean:
	rm ./main