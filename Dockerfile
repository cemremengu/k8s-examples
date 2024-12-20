FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/bin/ .

FROM alpine
RUN apk --no-cache add curl
COPY --from=builder /usr/bin/test /usr/bin/
ENTRYPOINT ["/usr/bin/test"]
