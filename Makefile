help:
	@echo -e "Available targets:\n\n   test\n   build\n   run\n   validate-schema"

.PHONY: test
test:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: build
build:
	go build .

.PHONY: run
run:
	go run . 1>/dev/null

.PHONY: validate-schema
validate-schema:
	yq -j config.yaml -o=config.json
