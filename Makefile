.PHONY: up
up:
	docker-compose up --build

.PHONY: down
down:
	docker-compose down --volumes

definition-api:
	go get -d -v ./...
	go install -v ./...
	go build -o definition-api ./cmd/definition-api

.PHONY: githooks
githooks:
	git config core.hooksPath .githooks
