FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/worker ./cmd/worker/main.go

# Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/api /bin/api
COPY --from=builder /bin/worker /bin/worker

COPY configs /app/configs

CMD ["/bin/api"]
# or worker
