# ---
LOCAL_DB_CONTAINER_NAME = postgres-go-timeseries
LOCAL_DB_PASSWORD = password

local-db:
	docker stop $(LOCAL_DB_CONTAINER_NAME) || true
	docker rm $(LOCAL_DB_CONTAINER_NAME) || true
	docker run --name $(LOCAL_DB_CONTAINER_NAME) -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres > /dev/null 2>&1 || docker start $(LOCAL_DB_CONTAINER_NAME)
	docker cp ./pkg/infrastructure/schema.sql $(LOCAL_DB_CONTAINER_NAME):/docker-entrypoint-initdb.d/schema.sql

# ---
DB_PACKAGE = ./pkg/infrastructure

install-sqlc:
	go get github.com/kyleconroy/sqlc/cmd/sqlc

generate-db-code: install-sqlc
	cd $(DB_PACKAGE); sqlc generate

# ---
tidy:
	go mod tidy

build: tidy
	mkdir -p build
	go build -o ./build/timeseries-poc ./cmd/timeseries-poc
	chmod +x ./build/timeseries-poc

GO_MAIN_FILE = ./cmd/timeseries-poc/main.go
CONFIG_FILE = ./properties.json
run: build
	CONFIG=$(CONFIG_FILE) ./build/timeseries-poc
