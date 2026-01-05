.PHONY: run
run:
	docker compose up

.PHONY: frontend_exec
frontend_exec:
	docker compose exec frontend bash

.PHONY: backend_exec
backend_exec:
	docker compose exec backend bash

.PHONY: setup_without_user
setup_without_user: setup_frontend

.PHONY: setup_with_user
setup_with_user:
	cp docker_dev/compose.override.yml .
	echo "host_user_name=${USER}" > .env
	echo "host_group_name=${USER}" >> .env
	echo "host_uid=`id -u`" >> .env
	echo "host_gid=`id -g`" >> .env
	${MAKE} setup_frontend

.PHONY: setup_frontend
setup_frontend:
	docker compose run --rm frontend pnpm install

.PHONY: generate_api_code
generate_api_code:
	docker run --rm -u $(shell id -u):$(shell id -g) -v ${PWD}/openapi:/spec redocly/cli bundle ./openapi.yaml -o ./openapi.gen.yaml
	docker compose run --rm backend go generate ./openapi
	docker compose run --rm frontend pnpm run gen:openapi

.PHONY: build_all
build_all:
	docker build -t go-react-app .

.PHONY: run_image
run_image:
	docker run -p 8080:8080 go-react-app

.PHONY: generate_ent
generate_ent:
	docker compose run --rm backend go generate ./internal/infra/ent

.PHONY: generate_migration_diff
generate_migration_diff:
	docker compose run --rm backend atlas migrate diff --env local

.PHONY: apply_migration
apply_migration:
	docker compose run --rm backend atlas migrate apply --env local

.PHONY: migration_status
migration_status:
	docker compose run --rm backend atlas migrate status --env local

.PHONY: seed
seed:
	docker compose run --rm backend go run ./cmd/seed/main.go
