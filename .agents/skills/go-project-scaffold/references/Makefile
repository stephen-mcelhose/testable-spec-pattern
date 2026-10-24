.PHONY: build test lint generate tidy

build:
	go build ./...

test:
	go test -race ./...

lint:
	golangci-lint run ./...

# Regenerate all counterfeiter fakes.
# Each interface that needs a fake should carry a //go:generate directive:
#   //go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Store
generate:
	go generate ./...

tidy:
	go mod tidy
