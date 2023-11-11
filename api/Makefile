DSN=host=localhost port=5432 user=postgres password=password dbname=api sslmode=disable timezone=UTC connect_timeout=5
BINARY_NAME=api

build:
	go build -o ${BINARY_NAME} ./cmd/api/

run: build
	@env DSN=${DSN} ./${BINARY_NAME}

clean:
	@go clean

start: run

stop:
	@-pkil -SIGTERM -f "./${BINARY_NAME}"