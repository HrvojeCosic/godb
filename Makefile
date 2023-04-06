.PHONY: test

test:
	go test ./src/...

lint:
	golangci-lint run -E gocritic