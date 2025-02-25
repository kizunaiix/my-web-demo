# GinDemo

用gin试着写一些api

## 项目结构

小型的go项目结构：

```text
gin/
├── cmd/                      # 入口文件 -> 可能有多个app
├── internal/                 # 业务逻辑
│    ├── handler/             # 控制器层
│    ├── service/             # 业务逻辑层
│    ├── repository/          # DAO 层
│    └── model/               # 数据模型
├── pkg/                      # 工具包
├── conf/                     # 配置文件
├── migrations/               # 数据库迁移脚本
└── main.go                   # 入口文件 -> 可能是用于dev
```

| 目录                          | 作用                                                | 细节                                                                |
| ----------------------------- | --------------------------------------------------- | ------------------------------------------------------------------- |
| handler/ (控制器层)           | 负责 解析 HTTP 请求和返回 JSON 响应，不写业务逻辑   | 解析 c.Param() / c.Query()，调用 service/ 处理业务逻辑              |
| service/ (业务逻辑层)         | 负责 业务逻辑，调用 repository/ 访问数据库          | 处理数据校验、事务、业务规则等                                      |
| repository/ (数据访问层，DAO) | 封装数据库操作，使用 GORM 或手写 SQL                | 避免在 service/ 直接写 SQL，保证数据库操作的可复用性                |
| model/ (数据模型层)           | 定义 struct，存储数据库表映射（ORM 模型）和业务模型 | 可以拆成 model/db.go（数据库模型）和 model/dto.go（前端返回的模型） |
| pkg/ (工具函数)               | 存放通用的工具函数或库，与业务无关                  | 如日志、时间转换、排序算法、JWT 解析等                              |
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
