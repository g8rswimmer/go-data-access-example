.PHONY: run_local
run_local:
	go run cmd/user-server/*.go

.PHONY: test
test:
	go test ./... -cover