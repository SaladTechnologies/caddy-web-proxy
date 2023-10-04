FROM golang:1.21-bookworm AS builder

WORKDIR /app
COPY . .

RUN go mod download && go mod vendor
RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
RUN xcaddy build --with github.com/SaladTechnologies/caddy-web-proxy

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/caddy .
COPY Caddyfile .

EXPOSE 3000

CMD ["./caddy", "run", "--config", "Caddyfile"]