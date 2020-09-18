help:  ## This Makefile's options.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\%-16s %s\n", $$1, $$2}'

test: ## Do standard code test.
	gofmt -s -l .
	goimports -l -format-only .
	go vet ./...
	go test -v -cover -race ./...