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

lint:
	golangci-lint run

down:
	kind delete cluster --name playground

up: 
	KUBECONFIG=~/.kube/playground kind create cluster --name=playground 

run:
	go run .
