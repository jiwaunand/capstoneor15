FROM golang:1.21-alpine AS builder
WORKDIR /app
RUN go mod init go-s
COPY main.go .
RUN go build -o server main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY public/ ./public/
EXPOSE 8080
CMD [ "./server" ]
