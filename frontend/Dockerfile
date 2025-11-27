FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -o app main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/app .
COPY templates ./templates
# The binary is static; we only need the binary and templates
CMD ["./app"]
