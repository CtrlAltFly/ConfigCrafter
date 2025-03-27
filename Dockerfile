FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./main.go

# Final stage
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]