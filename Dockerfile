FROM golang:latest

WORKDIR /app

COPY . .

ENV PORT=8080

EXPOSE ${PORT}

CMD [ "go" ,"run", "main.go" ]
