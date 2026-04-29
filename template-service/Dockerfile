FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy shared module first
COPY shared /app/shared

# Copy service files
COPY template-service/go.mod template-service/go.sum ./
COPY template-service/ ./

# Fix replace path to match container layout
RUN go mod edit -replace=github.com/sushantpardhi/shared=./shared

RUN go mod download
RUN CGO_ENABLED=0 go build -o template-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/template-service .
COPY template-service/.env.development ./.env.development
COPY template-service/.env.production ./.env.production

EXPOSE 8083
CMD ["./template-service"]