# Build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o wechat

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/wechat /bin/
ENTRYPOINT /bin/wechat
LABEL Name=drone-wechat-work Version=0.1.0
