# Stage 1: Build the Go application
FROM golang:1.21-alpine as builder

WORKDIR /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download

COPY ../../ ./

RUN go build -o ./bin/chatapi ./cmd/chatapi/main.go

# Stage 2: Create a minimal runtime container
FROM alpine

WORKDIR /app

COPY --from=builder /app/bin/chatapi .

CMD ["./chatapi"]