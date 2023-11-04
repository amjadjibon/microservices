auth-run:
	@echo "Running auth service..."
	@go run cmd/auth/main.go run

auth-migrate-up:
	@echo "Running auth service migration up..."
	@go run cmd/auth/main.go migrate up

auth-migrate-down:
	@echo "Running auth service migration down..."
	@go run cmd/auth/main.go migrate down
