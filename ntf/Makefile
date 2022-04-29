# Packages
PACKAGES=$(shell go list ./...|grep -v -e mocks|tr '\n' ' ')

# Command
download:
	@go mod download

test:
	$(foreach pkg, $(PACKAGES) ntf,\
    gotest -race -failfast -short -v $(pkg);)

cover:
	@echo "mode: cover" > coverage-data.txt
	$(foreach pkg, $(PACKAGES) ntf, \
	gotest -failfast -v -coverprofile=coverage.out -covermode=set $(pkg);\
	tail -n +2 coverage.out >> coverage-data.txt;)
	go tool cover -func=coverage-data.txt | grep total:
	rm coverage.out

run:
	@go run main.go

run-docker:
	@docker-compose up

stop-docker:
	@docker-compose down

mocks:
	@mockery --all --recursive --keeptree

.PHONY: download test cover run run-docker stop-docker mocks
