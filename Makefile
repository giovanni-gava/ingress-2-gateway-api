run:
	go run cmd/main.go

build:
	go build -o bin/gateway-migrator cmd/main.go

test:
	go test ./...

wire:
	cd internal/di && wire

clean:
	rm -rf bin
