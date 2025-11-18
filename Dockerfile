FROM golang:1.25-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags="-w -s" -o myapp .

FROM alpine:latest
RUN apk --no-cache add ca-certificates && \
    adduser -D -s /bin/sh appuser && \
    mkdir -p /app/data && \
    chown -R appuser:appuser /app

WORKDIR /app
COPY --from=builder /app/myapp .
RUN chmod 500 myapp

USER appuser
VOLUME ["/app/data"]
EXPOSE 8080
CMD ["./myapp"]
