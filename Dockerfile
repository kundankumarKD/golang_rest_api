# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o server cmd/api/main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

# Install dependencies for SQLite (if needed, though static build above uses CGO)
# Using alpine, we might need libc6-compat if not purely static
RUN apk --no-cache add ca-certificates sqlite

COPY --from=builder /app/server .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./server"]
