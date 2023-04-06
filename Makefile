.PHONY: test

test:
	go test ./test/...

lint:
	golangci-lint run -E gocritic