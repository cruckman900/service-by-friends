FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -o app main.go

FROM scratch
COPY --from=build /app/app /app
# The binary is static; we only need the binary
ENTRYPOINT ["./app"]
