# Build stage
FROM golang:1.24-alpine AS builder

RUN apk --no-cache add make

WORKDIR /app
COPY . .

RUN make release

FROM alpine:latest

RUN mkdir /etc/gateway
COPY --from=builder /app/build/ezex-gateway /usr/bin/ezex-gateway

EXPOSE 8080

ENTRYPOINT ["/usr/bin/ezex-gateway"]
