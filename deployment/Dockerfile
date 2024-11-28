FROM golang:1.23-alpine
WORKDIR /musicapp
COPY .. .

RUN go mod download

COPY .. .

RUN go build "./cmd/musicapp/main.go"
EXPOSE 8080
CMD ["./main"]