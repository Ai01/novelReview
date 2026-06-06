## Why

当前后端代码全部放在根目录 `package main` 下，所有文件平铺（`main.go`、`handlers.go`、`models.go`、`database.go`、`auth.go`），随功能迭代文件将持续膨胀，定位困难、耦合严重。需要按 MVC+DDD 分层重整代码结构——domain 层按领域实体（用户、书籍、评论、书单等）拆分，为后续书库、书单、个人中心等页面开发提供清晰的模块边界。

## What Changes

- **DDD 领域拆分**：domain 层按实体拆分为 `user/`、`book/`、`comment/` 子包，每个实体独立管理自己的模型和基础行为
- **目录结构重组**：按 MVC+DDD 分层，引入 `internal/` 目录，拆分为 domain（按实体分子包）、repository（按实体分文件）、service（业务逻辑）、handler（按实体分文件）、router（路由注册）、infra（基础设施）六个子包
- **代码迁移（零功能变更）**：将现有代码按职责迁移到对应目录，API 行为、数据库连接逻辑、JWT 逻辑完全不变
- **依赖注入**：`main.go` 作为组装入口，显式构建依赖链（infra → repository → service → handler → router）
- **更新 README.md**：反映新的目录结构和架构说明

## Capabilities

### New Capabilities
- `backend-layered-architecture`: 后端分层架构规范，定义 domain/repository/service/handler/router/infra 六层职责和依赖规则

### Modified Capabilities
<!-- 无 spec 变更，仅为代码重组，API 行为不变 -->

## Impact

- **后端全部文件**：`backend/*.go` 迁移重组（零业务逻辑变更）
- **Dockerfile**：`COPY . .` → 兼容 `internal/` 子目录
- **README.md**：更新架构描述和目录结构
- **deploy**：无变更（docker-compose、nginx、监控配置不变）
- **前端**：无变更
