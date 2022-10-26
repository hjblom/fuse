.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test ./internal/... -cover -race -mod=vendor -v

.PHONY: install
install:
	go install
