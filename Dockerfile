FROM golang:1.22.1-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/api/main.go

FROM alpine:3.19
WORKDIR app
COPY --from=builder /app .
CMD ["./app"]