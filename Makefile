.PHONY: generate

generate:
	@go generate ./...
	@echo "[OK] The static GraphQL files have been successfully generated"

