GORUN=go run
GOTEST=go test

run:
	$(GORUN) cmd/main.go

test:
	$(GOTEST) -v ./pkg/...

bench:
	$(GOTEST) -v -bench=. ./pkg/...