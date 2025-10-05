# stage 1 - build
FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/main-service ./cmd/main-service
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/worker-service ./cmd/worker-service

# stage 2 - final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /out/main-service .
COPY --from=builder /out/worker-service .