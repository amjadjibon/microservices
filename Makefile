auth-run:
	@echo "Running auth service..."
	ENV=dev go run cmd/auth/main.go run

auth-migrate-up:
	@echo "Running auth service migration up..."
	@ENV=dev go run cmd/auth/main.go migrate up

auth-migrate-down:
	@echo "Running auth service migration down..."
	@ENV=dev go run cmd/auth/main.go migrate down

title-run:
	@echo "Running title service..."
	ENV=dev go run cmd/title/main.go run

title-migrate-up:
	@echo "Running title service migration up..."
	@ENV=dev go run cmd/title/main.go migrate up

title-migrate-down:
	@echo "Running title service migration down..."
	@ENV=dev go run cmd/title/main.go migrate down
