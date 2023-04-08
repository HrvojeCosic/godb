.PHONY: test

test:
	go test ./...

lint:
	golangci-lint run -E gocritic