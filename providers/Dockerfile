FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]