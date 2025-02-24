# GinDemo

用gin试着写一些api

## SWAGGER相关

cd到根目录`cd /home/yaodong/codings/my-web-demo/gin`

在gin目录执行`swag init -d ./cmd,./handlers`，这样才会在gin下面生成doc目录

`go build -tags dev -o ./bin/gin-app-be ./cmd/`命令能够看到swagger文档，不带tags则不行。这里要注意是./cmd/ ，不能用./cmd/*.go或者./cmd/main.go
文件夹的话go会自动判断tag的逻辑。而指定go文件的话会高于tag逻辑，直接无视掉tag

### 注释的写法 


```go
// @Param [参数名] [参数类型] [数据类型] [是否必要：false|true] "描述"
```

其中：
- 参数类型有：query | path | header | body | formData
- 数据类型有：string (string) | integer (int, uint, uint32, uint64) | number (float32) | boolean (bool) | user defined struct

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
