stubs:
	protoc -I ./pb --go_out ./internal/api --go_opt paths=source_relative --go-grpc_out ./internal/api --go-grpc_opt paths=source_relative ./pb/ledger.proto
