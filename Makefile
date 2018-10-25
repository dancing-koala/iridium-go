GORUN=go run
GOTEST=go test

run:
	$(GORUN) cmd/main.go

test:
	$(GOTEST) -v ./pkg/...

bench:
	$(GOTEST) -v -bench=. ./pkg/...

test_lexer:
	$(GOTEST) -v ./pkg/iridium/assert.go ./pkg/iridium/lexer.go ./pkg/iridium/lexer_test.go