FROM golang:1.18 as build-env
# 开下梯子
ENV GOPROXY=https://goproxy.cn,direct

# 因为需要访问外网，所以在build 的时候需要配置网络 --network host
# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -gcflags="all=-N -l" -o server
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app server.go


# Final stage
#FROM debian:buster
FROM ubuntu

EXPOSE 8888 40000

WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /dockerdev/server /

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient","--check-go-version=false", "exec", "/server"]
#CMD ["/server"]