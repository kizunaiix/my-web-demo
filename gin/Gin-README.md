# GinDemo

用gin试着写一些api

## SWAGGER相关

cd到根目录`cd /home/yaodong/codings/my-web-demo/gin`

在gin目录执行`swag init -d ./cmd,./handlers`，这样才会在gin下面生成doc目录

## 构建命令

```zsh

# 项目构建
swag init -d ./cmd,./handlers
go build -o ./bin/gin-app-be ./cmd/*.go

# 镜像构建
docker build -t kizunaiix/my-web-demo-backend:latest .

# docker run
docker run -itd --name my-web-demo-backend -p 9000:9000 kizunaiix/my-web-demo-backend:latest
```
