FROM golang:1.23.4-alpine

WORKDIR /app

COPY . .

ENTRYPOINT [ "go run main.go" ]
