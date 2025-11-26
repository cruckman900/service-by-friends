FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/app .
CMD ["./app"]
