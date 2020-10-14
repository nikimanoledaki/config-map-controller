.PHONY: all test clean

test: 
	go test ./pkg/... 

race:
	go test -race ./pkg/... -short

coverage: 
	mkdir -p ./artifacts
	go test ./pkg/... -cover -coverprofile=./artifacts/coverage.out fmt
	go tool cover -html=artifacts/coverage.out -o artifacts/coverage.html
	rm ./artifacts/coverage.out
