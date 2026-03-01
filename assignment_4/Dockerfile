FROM golang:1.25.6-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api/main.go

FROM alpine

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]