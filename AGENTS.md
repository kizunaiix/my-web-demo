# AGENTS.md

给在本仓库工作的编码 agent 的约定。与用户交流用中文，提交信息也用中文。

## 项目概览

习惯养成/打卡应用，单仓前后端：
- `frontend/`：bun + React 19 + React Router 8（framework 模式，`ssr: false` 的 SPA，basename `/my-web-demo/`）+ Tailwind 4 + shadcn/ui。
- `backend/`：Go 1.25 + Gin + zap + swaggo；GORM/PostgreSQL 仅有接口占位，当前使用内存 repository。
- 部署：Gin 托管前端产物、nginx 反代，容器化/minikube 部分尚未验证。

## 常用命令

优先使用 Taskfile，不要另造命令：
- `task swag:init`：生成 Swagger 文档
- `task go:run-dev` / `task go:run-prod`：启动后端（自动依赖 swag:init）
- `task go:test`、`task go:build`（产物在 `backend/bin/`）
- `task bun:build`：前端构建并复制到 `backend/public/`

前端在 `frontend/` 下：`bun install --frozen-lockfile`、`bun run dev`、`bun run build`、`bun run typecheck`、`bun run format`。禁止使用 npm/yarn。

## 目录与后端分层

`backend/cmd/` 为入口；`backend/internal/<feature>/` 按 `handler.go`（绑定/校验/响应）、`service.go`（业务规则）、`repository.go`（存储）、`model.go`、`dto.go` 分层；错误统一用 `internal/middleware/myerr`，可复用代码放 `pkg/`，测试放 `backend/test/`。

- handler 不得写业务逻辑或直接操作存储，只能调用 service。
- 用 `ctx.ShouldBindJSON`，错误用 `ctx.Error(myerr.ErrXxx(...))` 交给中间件，不要用 `BindJSON`。
- 响应统一 `dto.UniResponseBody`，不要手拼非标准 JSON。
- 实现 repository 时保留接口断言 `var _ XxxRepository = (*impl)(nil)`。

## API 与 Swagger

- RPC 风格，只用 GET/POST，路径前缀 `/api/my-web-demo`，禁止 PUT/DELETE/PATCH。
- 每个接口写 swag 注释（`@Tags`、`@Router ... [post]`、`@Param`、`@Success`），改完执行 `task swag:init`。
- Swagger 仅在 `ENV` 不含 prod 时挂载于 `/swagger/index.html`。

## 前端约定

- 路由统一在 `frontend/app/routes.ts` 注册，路由文件放 `app/routes/`，用 kebab-case。
- shadcn 组件放 `app/components/ui/` 并保持 kebab-case 文件名，不要重命名以免 CLI 失效；业务组件用 PascalCase。
- 用 `~/` 别名导入 `app/` 下代码；样式只用 Tailwind，条件类名用 `cn()`。
- `react-router.config.ts` 的 `ssr: false` 与 `basename` 不要随意改动。

## 禁止事项

- 不要手改生成物：`frontend/build/`、`backend/public/`、`backend/docs/`、`bun.lock`（除非确需变更依赖）。
- 不要提交 `node_modules/`、构建产物或本地配置；不要用 `panic` 处理可恢复错误。
- 不要改动 `frontend/old/`、`frontend/.next-deprecated/` 等遗留目录，除非明确要求。

## 已知坑

- 后端必须在 `backend/work/` 下运行（Taskfile 已处理），否则 `../public` 静态资源和 `NoRoute` 回落会失效。
- 前端构建产物固定为 `frontend/build/client`，由 `task bun:build` 复制到 `backend/public/`。
- 数据库尚未接入；新增持久化需求先补 repository 实现，不要在 handler 里直连 DB。
- Docker/minikube/nginx 相关任务仍是 TODO，改动前先确认。

## 提交前检查

`task go:test`、`cd frontend && bun run typecheck`、`bun run format` 全部通过后再提交；提交信息用中文，风格与现有历史一致。
