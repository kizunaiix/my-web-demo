# GinDemo

用gin试着写一些api

## 项目结构

小型的go项目结构：

```text

my-project-gin/
├── cmd/
│   └── main.go
├── internal/
│   ├── router/
│   │   └── router.go         # 注册所有模块路由
│   ├── server/
│   │   └── run.go            # 组装依赖 & 启动 HTTP 服务    #TODO 加上这个 
│   ├── {业务模块}/
│   │   ├── handler.go        # 控制器
│   │   ├── service.go        # 业务逻辑
│   │   ├── repository.go     # DAO
│   │   ├── model.go          # 数据模型 (对应数据库的表结构)
│   │   └── dto.go            # 数据传输对象 (例如api的请求体和响应体)
│   ├── .../
│   └── .../
├── pkg/
└── conf/
    └── config.yml

```

| 文件                          | 作用                                                | 细节                                                                |
| ----------------------------- | --------------------------------------------------- | ------------------------------------------------------------------- |
| handler.go                      | 负责 解析 HTTP 请求和返回 JSON 响应，不写业务逻辑   | 解析 c.Param() / c.Query()，调用 service/ 处理业务逻辑              |
| service.go                   | 负责 业务逻辑，调用 repository/ 访问数据库          | 处理数据校验、事务、业务规则等                                      |
| repository.go ，也有叫dao.go    | 封装数据库操作，使用 GORM 或手写 SQL                | 避免在 service/ 直接写 SQL，保证数据库操作的可复用性                |
| model.go                      | 定义 struct，存储数据库表映射（ORM 模型）和业务模型 | 可以拆成 model/db.go（数据库模型）和 model/dto.go（前端返回的模型） |
| pkg/ (工具函数)                 | 存放通用的工具函数或库，与业务无关                  | 如日志、时间转换、排序算法、JWT 解析等                              |

## SWAGGER相关

cd到cmd目录,执行`swag init -d ./,../internal/handler`，swag会在当前cd到的目录下面生成doc目录

~~`go build -tags dev -o ./bin/gin-app-be ./cmd/`命令能够看到swagger文档，不带tags则不行。这里要注意是./cmd/ ，不能用./cmd/*.go或者./cmd/main.go  .  文件夹的话go会自动判断tag的逻辑。而指定go文件的话会高于tag逻辑，直接无视掉tag~~

现在会根据环境变量ENV里是否包含"prod"来保证不在生产环境加载swagger文档。

### 注释的写法

```go
// @Param <参数名> <位置> <数据类型> <是否必要：false|true> "描述"
```

其中：

- 位置有：query | path | header | body | formData
- 数据类型有：string (string) | integer (int, uint, uint32, uint64) | number (float32) | boolean (bool) | user defined struct

## 构建命令

~~参考makefile，命令为`make <command>`~~

---

改用了Task

- 初始化并运行  
`task gorun`

- 编译 Go 二进制  
`task gobuild`

- 构建 Docker 镜像  
`task dockerbuild`

- 启动开发容器  
`task dockerrun-dev`

- 删除容器（可选）  
`task docker-clean`
