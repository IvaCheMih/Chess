.PHONY: start
.PHONY: all test clean

lint:
	golangci-lint run --fix

vuln:
	govulncheck ./...

test:
	go clean -testcache
	go test -v ./... -short

run-local:
	go run src/main.go
