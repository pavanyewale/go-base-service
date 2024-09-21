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

install: 
	go mod tidy
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

setup:
	buf dep update
	buf generate ./proto/v1/*.proto
	go mod tidy