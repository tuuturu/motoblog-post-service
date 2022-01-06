
fmt:
	@goimports -w .
	@gofmt -w .

test:
	@go test ./...

run:
	@go run main.go
