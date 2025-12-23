FROM golang:1.23 AS builder

WORKDIR /app

# Copy entire monorepo into the build context
COPY . .

# Move into the service directory
WORKDIR /app/services/auth-service

# Download dependencies
RUN go mod download

# Build the service
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# ---- Run Stage ----
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/services/auth-service/server .

EXPOSE 8001

CMD ["./server"]
