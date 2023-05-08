FROM golang:1.17.1-alpine as builder

WORKDIR /auth
COPY go.mod go.sum ./
RUN go mod download
COPY .. ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o /api -ldflags '-s -w' ./cmd/api

FROM alpine:3.15

# アプリケーションを実行するユーザーを作成
RUN addgroup -S group && adduser -S user -G group

# ビルドしたバイナリをコピー
COPY --from=builder /api /usr/local/bin/api
COPY --from=builder /auth/config/config.json /usr/local/bin/config.json

# アプリケーションを実行するユーザーに切り替え
USER user

ENTRYPOINT ["/usr/local/bin/api"]
