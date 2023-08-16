COVERAGE_FILE = coveragefile

generate:
	@echo "generating graphql models..."
	@go generate ./...

test-cover:
	go test -v ./... -coverprofile=${COVERAGE_FILE}
	go tool cover -html=${COVERAGE_FILE} && go tool cover -func ${COVERAGE_FILE} && unlink ${COVERAGE_FILE}