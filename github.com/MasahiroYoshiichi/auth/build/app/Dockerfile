# ビルドステージ
FROM node:14 AS builder

WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn build

# 本番ステージ
FROM nginx:stable-alpine

# Reactアプリケーションのビルドファイルをコピー
COPY --from=builder /app/build /usr/share/nginx/html

# Nginxの設定ファイルをコピー
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
