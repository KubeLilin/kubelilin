# 构建容器,使用版本: golang:1.16
FROM golang:1.16-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /publish

COPY . .

RUN go build -ldflags="-s -w" -o app .

# 最终容器
FROM scratch:latest as production

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=build /publish/app /

COPY --from=build /publish/*.yml /

# Run
CMD ["/app"]