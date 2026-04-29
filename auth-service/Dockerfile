FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy shared module first
COPY shared /app/shared

# Copy service files
COPY auth-service/go.mod auth-service/go.sum ./
COPY auth-service/ ./

# Fix replace path to match container layout
RUN go mod edit -replace=github.com/sushantpardhi/shared=./shared

RUN go mod download
RUN CGO_ENABLED=0 go build -o auth-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/auth-service .
COPY auth-service/.env.development ./.env.development
COPY auth-service/.env.production ./.env.production

EXPOSE 8001
CMD ["./auth-service"]