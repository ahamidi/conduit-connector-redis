.PHONY: build test

build:
	go build -o conduit-connector-redis cmd/redis/main.go

test:
	# run required docker containers, execute integration tests, stop containers after tests
	docker compose -f test/docker-compose-redis.yml up --quiet-pull -d --wait
	go test $(GOTEST_FLAGS) -v -race ./...; ret=$$?; \
		docker compose -f test/docker-compose-redis.yml down; \
		exit $$ret