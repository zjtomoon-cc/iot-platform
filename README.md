# iot-platform

> 基于 go-zero(微服务) 实现的物联网平台
> 
> B站视频地址：https://www.bilibili.com/video/BV13G4y1R71m

## 技术栈

+ 后端：go-zero、gorm
+ 前端：vue
+ 硬件：arduino、esp8266

## 安装

1. 搭建 [Golang](https://golang.google.cn/) 环境
2. 安装 goctl
```shell
go get -u github.com/zeromicro/go-zero/tools/goctl@latest 
```
3. 下载安装 [arduino](https://www.arduino.cc/en/donate/)
4. 搭建 ETCD 环境
```shell
# 参考如下，使用Docker安装ETCD
docker run -d --name Etcd-server \
    --network app-tier \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
```

## 命令

+ 创建API服务

```shell
goctl api new 服务名称
# 1. 创建 user 服务
goctl api new user
# 2. 创建 admin 服务
goctl api new admin
```

+ 生成服务代码

```shell
goctl api go -api 服务名称.api -dir . -style go_zero
# 1. 生成 user api 服务代码
goctl api go -api user.api -dir . -style go_zero
# 2. 生成 admin api 服务代码
goctl api go -api admin.api -dir . -style go_zero
# 3. 生成 user rpc 服务代码
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
```

+ 启动服务

```shell
go run 服务名称.go -f 配置文件地址
# 1. 启动 user api 服务
go run user.go -f etc/user-api.yaml
# 2. 启动 admin api 服务
go run admin.go -f etc/admin-api.yaml
# 3. 启动 user rpc 服务
go run user.go -f etc/user.yaml
```

## 适用场景

共享单车、共享充电宝、外卖柜

## 功能模块

+ [ ] 用户模块
  + [x] 登录
+ [ ] 后台管理模块
  + [ ] 设备管理
    + [x] 设备列表
  + [ ] 产品管理
    + [x] 产品列表
    + [x] 创建、修改产品
+ [ ] 开放平台模块
+ [ ] 设备服务模块
