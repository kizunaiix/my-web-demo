# my-web-demo

学习web技术，前端是基于bun的SPA应用，后端是gin，gin上挂前端产物以及后端api，整个应用做容器化，上minikube部署，最前面挂nginx代理

目标是写一个习惯养成/打卡应用，技术栈包括：

    前端：TypeScript：ReactRouter框架模式、shadcn/ui
    后端：Go：Gin
    数据库：PostgreSQL
    API 设计：偏向 RPC 风格，只使用 GET 和 POST 方法，不使用 PUT 和 DELETE
    Swagger 文档：swaggo

## gin

[README](./backend/Gin-README.md)
