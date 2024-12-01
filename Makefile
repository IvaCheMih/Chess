.PHONY: start
.PHONY: all test clean

lint:
	golangci-lint run --fix

vuln:
	govulncheck ./...

test:
	go clean -testcache
	go test -v ./... -short

run-remote:
	docker-compose -f docker-compose-remote.yaml up -d

run-local-all:
	docker-compose -f docker-compose-local.yaml up -d
	sleep 2
	go run src/main.go

run-local-docker:
	docker-compose -f docker-compose-local.yaml up -d

run-local-app:
	go run src/main.go

swagger:
	( cd src ; swag init )