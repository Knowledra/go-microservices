FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy shared module first
COPY shared /app/shared

# Copy service files
COPY user-service/go.mod user-service/go.sum ./
COPY user-service/ ./

# Fix replace path to match container layout
RUN go mod edit -replace=github.com/sushantpardhi/shared=./shared

RUN go mod download
RUN CGO_ENABLED=0 go build -o user-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-service .
# Copy the service .env file for godotenv
COPY user-service/.env .env

EXPOSE 8002
CMD ["./user-service"]