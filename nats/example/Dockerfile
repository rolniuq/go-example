FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o nats-example .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/nats-example .

CMD ["./nats-example"]
