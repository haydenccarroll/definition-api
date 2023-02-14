.PHONY: up
up:
	docker-compose up --build

.PHONY: down
down:
	docker-compose down --volumes

definition_api:
	go get -d -v ./...
	go install -v ./...
	go build -o definition_api ./cmd/definition_api

.PHONY: githooks
githooks:
	git config core.hooksPath .githooks
