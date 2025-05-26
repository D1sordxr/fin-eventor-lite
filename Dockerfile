FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/worker ./cmd/worker

# Final image
FROM alpine:latest
COPY --from=builder /bin/api /bin/api
COPY --from=builder /bin/worker /bin/worker

CMD ["/bin/api"]
# or worker
