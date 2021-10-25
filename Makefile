.PHONY: build
build:
	go build cmd/wrk-internship-api/main.go

.PHONY: test
test:
	go test -v ./...