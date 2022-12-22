.PHONY: all test clean build

test: lint
	ginkgo -r

e2e: lint build
	go test -timeout=30m ./test/acceptance -v

race: lint
	go test -race ./pkg/... -short

coverage: lint
	mkdir -p ./artifacts
	go test ./pkg/... -cover -coverprofile=./artifacts/coverage.out fmt
	go tool cover -html=artifacts/coverage.out -o artifacts/coverage.html
	rm ./artifacts/coverage.out

lint:
	golangci-lint run

build:
	go build -o build/controller main.go

down:
	kind delete cluster --name playground

up: 
	KUBECONFIG=~/.kube/playground kind create cluster --name=playground 

run:
	go run .

watch: 
	KUBECONFIG=~/.kube/playground kubectl get cm foo -oyaml --watch

.PHONY: prepare-release
prepare-release: ## Create release
	build/tag-release.sh
