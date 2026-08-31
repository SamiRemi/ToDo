.PHONY: test run clean

test:
	go test ./...

run:
	go run cmd/server/main.go

clean:
	rm -f ToDo-api