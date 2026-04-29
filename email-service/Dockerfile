FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy shared module first
COPY shared /app/shared

# Copy service files
COPY email-service/go.mod email-service/go.sum ./
COPY email-service/ ./

# Fix replace path to match container layout
RUN go mod edit -replace=github.com/sushantpardhi/shared=./shared

RUN go mod download
RUN CGO_ENABLED=0 go build -o email-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/email-service .
COPY email-service/.env.development ./.env.development
COPY email-service/.env.production ./.env.production

EXPOSE 8082
CMD ["./email-service"]