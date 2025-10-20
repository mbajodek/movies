.PHONY: build
build:
	go build -o ./bin/movies/main ./cmd/movies/main.go

.PHONY: run
run:
	go run ./cmd/movies/main.go

.PHONY: unit-test
unit-test:
	go test -v ./internal/...

.PHONY: k6-test
k6-test:
	k6 run internal/k6/post_movie_test.js

.PHONY: docker
docker:
	docker build -t go-movies:latest .

.PHONY: docker-compose
docker-compose:
	docker compose -f docker-compose.yml up -d

.PHONY: docker-compose-down
docker-compose-down:
	docker compose -f docker-compose.yml down
