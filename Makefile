default: debug test run

debug:
	go vet ./...
	gofmt -d ./
	gofmt -w ./
	go build -mod vendor -o ./out

test:
	go test ./...

run:
	@./out

update:
	go mod vendor
	go mod download
