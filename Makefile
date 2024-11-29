.PHONY: all test clean

lint:
	golangci-lint run --fix

vuln:
	govulncheck ./...

test:
	go test -v ./... -short