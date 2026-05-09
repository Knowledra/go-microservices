FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy shared module first
COPY shared /app/shared

# Copy service files
COPY department-service/go.mod department-service/go.sum ./
COPY department-service/ ./

# Fix replace path to match container layout
RUN go mod edit -replace=github.com/sushantpardhi/shared=./shared

RUN go mod download
RUN CGO_ENABLED=0 go build -o department-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/department-service .
COPY department-service/.env.development ./.env.development
COPY department-service/.env.production ./.env.production

EXPOSE 8001
CMD ["./department-service"]