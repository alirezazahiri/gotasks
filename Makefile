.PHONY: help migrate-up migrate-down migrate-version migrate-force migrate-create migrate-steps

help:
	@echo "Available migration commands:"
	@echo "  make migrate-up            - Run all pending migrations"
	@echo "  make migrate-down          - Rollback all migrations"
	@echo "  make migrate-version       - Show current migration version"
	@echo "  make migrate-steps N=1     - Run N migration steps (use negative for rollback)"
	@echo "  make migrate-force V=1     - Force database to specific version"
	@echo "  make migrate-create NAME=name - Create new migration files"

migrate-up:
	@echo "Running migrations up..."
	@go run cmd/migrate/main.go -cmd=up

migrate-down:
	@echo "Rolling back migrations..."
	@go run cmd/migrate/main.go -cmd=down

migrate-version:
	@go run cmd/migrate/main.go -cmd=version

migrate-steps:
	@go run cmd/migrate/main.go -cmd=steps -steps=$(N)

migrate-force:
	@go run cmd/migrate/main.go -cmd=force -version=$(V)

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=create_users_table"; \
		exit 1; \
	fi
	@TIMESTAMP=$$(date +%s); \
	UP_FILE="migrations/$${TIMESTAMP}_$(NAME).up.sql"; \
	DOWN_FILE="migrations/$${TIMESTAMP}_$(NAME).down.sql"; \
	touch $$UP_FILE $$DOWN_FILE; \
	echo "Created migration files:"; \
	echo "  $$UP_FILE"; \
	echo "  $$DOWN_FILE"

