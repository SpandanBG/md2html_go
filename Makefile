test:
	go test ./... -v -coverprofile=coverage.out

fmt:
	go fmt ./...
