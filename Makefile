# Путь к спецификации и конфигурации
OPENAPI_SPEC=./openAPI/openapi.yaml
OPENAPI_CONFIG=./openAPI/oapi-codegen.yaml
OPENAPI_OUTPUT=./internal/ogenerated/openapi.gen.go
# Миграции
DB_HOST=localhost
DB_PORT=5432
DB_USER=notif_user
DB_PASS=notif_pass
DB_NAME=notifications_db
MIGRATIONS_DIR=./migrations

DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose down

migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose version

migrate-force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose force 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Команда генерации
gen:
	oapi-codegen -config ./openAPI/oapi-config.yaml -include-tags notification -package notification ./openAPI/openapi.yaml > ./internal/ogenerated/api.gen.go

clgen:
	del /Q .\internal\ogenerated\api.gen.go

dc_up:
	docker compose up -d
dc_down:
	docker compose down

k6-getl:
	docker run -v "F:\test-REST_API:/project" -i grafana/k6 run /project/tests/k6/getl.js
k6-post:
	docker run -v "F:\test-REST_API:/project" -i grafana/k6 run /project/tests/k6/post.js
k6-ramp:
	docker run -v "F:\test-REST_API:/project" -i grafana/k6 run /project/tests/k6/ramp-test.js
run:
	go run ./cmd/app/main.go	

# Полезные команды
.PHONY: generate-openapi clean-openapi migrate-up migrate-down migrate-version migrate-force migrate-create clgen dc_up dc_down k6-getl k6-post