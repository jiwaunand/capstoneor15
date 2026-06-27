FROM golang:1.21-alpine AS builder
WORKDIR /app
RUN go mod init go-s
COPY main.go .
RUN go build -o server main.go

FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --from=builder /app/server .
COPY public/ ./public/
RUN chown -R appuser:appgroup /app
USER appuser
EXPOSE 8080
ENTRYPOINT ["./server"]
