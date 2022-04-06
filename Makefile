BINARY_NAME = fire
TEST_COMMAND = go test

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)

.PHONY: run
run: build
	./$(BINARY_NAME) 

.PHONY: test
test: 
	$(TEST_COMMAND) -cover -parallel 5 -failfast  ./... 

.PHONY: tidy
tidy:
	go mod tidy

# auto restart bot (using fiber CLI <3)
.PHONY: dev
dev:
	fiber dev -t ./cmd/$(BINARY_NAME)

.PHONY: lint
lint:
	revive -formatter friendly -config revive.toml ./...

.PHONY: spell
spell:
	misspell -error ./**

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: gosec
gosec:
	gosec -tests ./...

.PHONY: inspect
inspect: lint spell gosec staticcheck
