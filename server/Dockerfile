FROM node:18-alpine AS fe_build

WORKDIR /app

COPY ./client /app
# 安装 pnpm
RUN npm i -g pnpm --registry=https://registry.npmmirror.com
RUN pnpm config set registry https://registry.npmmirror.com
# 安装依赖
RUN pnpm install --frozen-lockfile
# 打包 cms-ui
RUN npm run build:fe

# 服务端打包
FROM golang:latest AS go_build

WORKDIR /app

COPY ./server /app

COPY --from=fe_build /app/apps/fe/dist /app/assets/fe

ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o opd main.go

# 服务运行
FROM alpine:latest

RUN apk update && apk add tzdata

WORKDIR /app

COPY --from=go_build /app/opd /app

VOLUME ["/apps/conf"]

ENV TZ="Asia/Shanghai"


EXPOSE 4002

# 启动
CMD ["/app/opd", "-c", "./config/config.prod.toml"]