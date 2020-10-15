.PHONY: all test clean

test: lint
	ginkgo -r

race:
	go test -race ./pkg/... -short

coverage: 
	mkdir -p ./artifacts
	go test ./pkg/... -cover -coverprofile=./artifacts/coverage.out fmt
	go tool cover -html=artifacts/coverage.out -o artifacts/coverage.html
	rm ./artifacts/coverage.out

build:
	go build -o build/controller .

lint:
	golangci-lint run
