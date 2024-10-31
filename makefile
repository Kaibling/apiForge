lint:
	golangci-lint run
pkg-update:
	go get -u
	go mod tidy
deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0