include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations


.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir ${MIGRATIONS_PATH} ${filter-out $@, $(MAKECMDGOALS)}

.PHONY: migrate-up
migrate-up:
	@migrate -path=${MIGRATIONS_PATH} -database=${DB_MIGRATOR_ADDR} up

.PHONY: migrate-down
migrate-down:
	@migrate -path=${MIGRATIONS_PATH} -database=${DB_MIGRATOR_ADDR} down ${filter-out $@, $(MAKECMDGOALS)}

.PHONY: migrate-force
migrate-force:
	@migrate -path=${MIGRATIONS_PATH} -database=${DB_MIGRATOR_ADDR} force ${filter-out $@, $(MAKECMDGOALS)}

.PHONY: migrate-delete-latest
migrate-delete-latest:
	@LATEST=$$(ls -1 ${MIGRATIONS_PATH}/*.up.sql | sort | tail -n 1); \
	BASENAME=$$(basename $$LATEST .up.sql); \
	rm -v ${MIGRATIONS_PATH}/$$BASENAME.up.sql ${MIGRATIONS_PATH}/$$BASENAME.down.sql

%:
	@:



