## Context

当前后端所有代码在 `backend/` 根目录下，`package main`，文件平铺。现有 5 个文件：`main.go`（入口+路由）、`handlers.go`（HTTP handler+业务逻辑）、`models.go`（GORM 模型）、`database.go`（数据库连接）、`auth.go`（JWT）。随着探索页上线，后续还有书库、书单、个人中心等多个模块，平铺结构难以维护。

## Goals / Non-Goals

**Goals:**
- 按 MVC+DDD 分层重组代码，形成清晰模块边界
- **DDD 领域拆分**：domain 层按实体（用户、书籍、评论）拆分子包，每个实体独立管理模型
- 使用 Go 标准 `internal/` 目录约束外部不可导入内部包
- `main.go` 成为纯组装入口，依赖显式构造
- Dockerfile 构建兼容新目录结构
- README.md 反映新架构

**Non-Goals:**
- 不修改任何业务逻辑、API 契约、数据库连接方式
- 不引入新依赖（不新增 DI 框架如 wire）
- 不拆分 go module（保持单 module）
- 不涉及前端和部署脚本变更

## Decisions

### 1. 分层结构：六层，domain 按实体拆分

```
backend/
├── main.go                          # 入口：组装依赖链，启动服务
├── internal/
│   ├── domain/                       # 领域模型层（按实体拆分子包，零依赖）
│   │   ├── user/
│   │   │   └── user.go              # User 结构体 (GORM 模型)
│   │   ├── book/
│   │   │   └── book.go              # Book 结构体 (GORM 模型)
│   │   └── comment/
│   │       └── comment.go           # Comment 结构体 (GORM 模型，含关联)
│   ├── repository/                  # 数据访问层（GORM 查询封装）
│   │   ├── user_repo.go             # 用户查询
│   │   └── comment_repo.go          # 评论查询 + 搜索
│   ├── service/                     # 业务逻辑层（编排 repository + 外部依赖）
│   │   ├── auth_service.go          # 登录逻辑 + JWT 签发
│   │   └── explore_service.go       # 探索/搜索逻辑
│   ├── handler/                     # HTTP 处理层（按实体拆分，参数解析、响应封装）
│   │   ├── auth_handler.go          # POST /api/v1/login
│   │   └── explore_handler.go       # GET /api/v1/explore, /api/v1/explore/search
│   ├── router/                      # 路由注册
│   │   └── router.go                # 集中注册所有路由
│   └── infra/                       # 基础设施
│       └── database.go              # 数据库连接、读写分离
```

**依赖方向（单向）：** `main → router → handler → service → repository → domain`，`infra` 由 `main` 注入给各层。

**domain 拆分原则**：每个领域实体独立一个子包。当前有 User、Book、Comment 三个实体，后续新增书单（BookList）时只需创建 `domain/booklist/` 子包，不影响已有代码。

**选择理由**：每层职责单一，新功能开发只需在对应层追加，不会影响其他层。领域拆分使每个实体的模型定义独立演进。

### 2. 为什么不引入 DI 框架（wire）

当前依赖链简单（一个数据库连接 + 几个 service），手写 `main.go` 组装更直观。数量可控时手写 DI 比 code generation 更易调试。后续服务膨胀到 10+ 时可再引入 wire。

### 3. domain 放在 internal/ 内

Go 的 `internal/` 目录由编译器强制拒绝外部 module 的 import，确保领域模型不会被外部项目意外依赖。

### 4. repository 层抽象

每个 repository 定义为接口 + 实现结构体：

```go
type CommentRepository interface {
    FindList(cursor string, limit int) ([]domain.Comment, bool, error)
    Search(q string, cursor string, limit int) ([]domain.Comment, bool, error)
}
```

Service 依赖接口，不依赖具体实现，便于后续单测 mock。

### 5. Dockerfile 适配

当前 Dockerfile 是 `COPY . .` + `go build -o main .`。迁移后 `main.go` 在根目录、其余在 `internal/`，构建命令不变：`go build -o main .` 仍能正确编译。

### 6. README.md 更新

新增"后端架构"章节，说明分层结构和各层职责。更新目录树。不删减现有的部署/配置/端口说明。

## Risks / Trade-offs

- **[风险] 存量代码迁移遗漏** → 分文件逐个迁移，每迁移一个跑 `go build` 验证
- **[风险] import 路径变化导致构建失败** → module path 不变（`github.com/user/novel-backend`），仅新增 `internal/xxx` 子包 import，Go 编译器会自动处理
- **[权衡] 文件数量从 5 增加到约 13** → 每个文件职责明确，定位更快。小型项目可能觉得过度设计，但后续 3+ 页面迭代会迅速体现价值

## Open Questions

- 无
