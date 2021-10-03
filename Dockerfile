# Deps stage
FROM golang:alpine3.14 AS deps
ADD go.mod /app/go.mod
ADD go.sum /app/go.sum
WORKDIR /app
RUN go mod download

# Build stage
FROM golang:alpine3.14 AS builder
ADD . /app
COPY --from=deps /go /go
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/main /app/cmd/main.go
RUN ls -la /build

# Final stage
FROM golang:alpine3.14
RUN apk --no-cache add  curl ca-certificates tzdata
COPY --from=builder /app/config /config
COPY --from=builder /app/docs /app/docs
COPY --from=builder /build /app
RUN ls
RUN chmod +x /app/main
ENTRYPOINT ["/app/main"]
EXPOSE 8080