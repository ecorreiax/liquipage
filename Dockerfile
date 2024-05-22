FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go build -o liquipage main.go

ENTRYPOINT ["./liquipage"]