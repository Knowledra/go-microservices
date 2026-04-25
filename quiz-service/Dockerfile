FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy shared module first
COPY shared /app/shared

# Copy service files
COPY auth-service/go.mod quiz-service/go.sum ./
COPY quiz-service/ ./

# Fix replace path to match container layout
RUN go mod edit -replace=github.com/sushantpardhi/shared=./shared

RUN go mod download
RUN CGO_ENABLED=0 go build -o quiz-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/quiz-service .

EXPOSE 8001
CMD ["./quiz-service"]