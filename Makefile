.PHONY: run docker

run:
	go run cmd/main.go

docker:
	cd docker && docker compose up -d