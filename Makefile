bench:
	go test -bench=. ./...

test:
	go test ./... -v -coverprofile=coverage.out

fmt:
	go fmt ./...
