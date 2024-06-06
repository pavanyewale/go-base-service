build: clean
	go build -o application .
build-linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application .
run:
	go run main.go
test:
	go test -v ./...
clean:
	rm -f application

docker: build-linux
	docker build -t go-base-service .
	rm -f application
docker-run:
	docker run -p 8080:8080 go-base-service

setup:
	go mod tidy