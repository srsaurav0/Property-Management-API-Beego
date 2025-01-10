# Build stage
FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/conf ./conf

RUN adduser -D appuser && \
    chown -R appuser:appuser /app

USER appuser
EXPOSE 8080

CMD ["./main"]