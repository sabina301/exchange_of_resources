FROM golang:1.23.0 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api_gateway .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api_gateway .

CMD ["./api_gateway"]


