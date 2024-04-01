.PHONY: dev
dev: fmt lint test

.PHONY: deps
deps:
	go install github.com/vektra/mockery/v2@v2.12.3

.PHONY: generate
generate: deps
	go generate $$(go list ./... | grep -v vendor)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover ./...

.PHONY: oas_test
oas_test:
	go test -v ./test -count=1 -run "TestOpenAPI"

.PHONY: bench
bench:
	go test -run XXXX -bench . -benchmem ./...
