# NovelWeb

## 架构

- 网关：Nginx（统一入口，转发前端与后端）
- 后端：Gin + Gorm（服务进程只提供业务能力，不做环境准备）
- 前端：Vite（容器内运行 dev server）
- 数据库：MySQL 主从（主写，从读；读库不可用时回退主库读）
- 监控：Prometheus（采集）+ Grafana（可视化）
- 日志：Loki（存储/查询）+ Promtail（采集）+ Grafana（查询入口）

## 后端架构（MVC + DDD 分层）

后端采用六层分层架构，依赖方向单向：`main → router → handler → service → repository → domain`

```
backend/
├── main.go                          # 组装入口（依赖注入，启动服务）
├── Dockerfile
├── go.mod / go.sum
└── internal/
    ├── domain/                       # 领域模型层（按实体拆分子包，零依赖）
    │   ├── user/user.go              # User 结构体（GORM 模型）
    │   ├── book/book.go              # Book 结构体（GORM 模型）
    │   └── comment/comment.go        # Comment 结构体（关联 User、Book）
    ├── repository/                   # 数据访问层（纯 GORM，接口+实现）
    │   ├── user_repo.go              # UserRepository 接口+查询用户
    │   └── comment_repo.go           # CommentRepository 接口+评论查询/搜索
    ├── service/                      # 业务逻辑层（编排 repository）
    │   ├── auth_service.go           # 认证（JWT 签发/校验 + 登录逻辑）
    │   └── explore_service.go        # 探索页（游标分页 + 搜索）
    ├── handler/                      # HTTP 处理层（参数解析、响应封装）
    │   ├── auth_handler.go           # POST /api/v1/login
    │   └── explore_handler.go        # GET /api/v1/explore, /api/v1/explore/search
    ├── router/router.go              # 路由集中注册
    └── infra/database.go             # 基础设施（数据库连接、读写分离）
```

| 层级 | 职责 | 依赖方向 |
|------|------|----------|
| **domain** | GORM 模型，纯数据结构 | 零依赖 |
| **repository** | 纯 GORM 数据查询，只依赖 domain | domain → |
| **service** | 业务逻辑编排，依赖 repository 接口 | repository → |
| **handler** | HTTP 参数解析/校验/响应，依赖 service | service → |
| **router** | 集中注册路由，依赖 handler | handler → |
| **main** | 组装入口，依赖注入链 | 以上全部 |

**变更新规则**：新增功能只需在对应层追加文件，不修改已有代码。

## 目录结构

- deploy/：仓库根目录第一级，部署与运维入口（docker compose、配置、脚本）
- backend/：后端服务（MVC+DDD 分层，internal/ 六层架构）
- frontend/：前端应用

## 端口与入口

- 网关（Nginx）：http://localhost:${GATEWAY_PORT:-80}
  - /healthz：网关健康检查
  - /api/*：转发到后端
  - /metrics：转发到后端 metrics
- 说明：Nginx 配置中的 `backend:8080` / `frontend:5173` 是容器网络内的服务端口，不受宿主机端口映射（BACKEND_PORT/FRONTEND_PORT）影响
- 后端：http://localhost:${BACKEND_PORT:-18080}
  - /healthz：存活检查
  - /readyz：就绪检查（至少依赖写库连通性）
  - /metrics：Prometheus 指标
  - /api/v1/login：登录接口
- 前端：http://localhost:${FRONTEND_PORT:-5173}（直连前端容器；推荐通过网关访问 http://localhost:${GATEWAY_PORT:-80}）
- Prometheus：http://localhost:${PROMETHEUS_PORT:-19090}
- Grafana：http://localhost:${GRAFANA_PORT:-3000}
- Loki：http://localhost:${LOKI_PORT:-3100}
- MySQL Master：localhost:${MYSQL_MASTER_PORT:-33306}
- MySQL Replica：localhost:${MYSQL_REPLICA_PORT:-33307}

## 配置约定（环境变量）

后端：
- WRITE_DSN：写库 DSN（指向 mysql-master）
- READ_DSN：读库 DSN（指向 mysql-replica，可缺省，缺省回退为 WRITE_DSN）
- JWT_SECRET：JWT 密钥

前端：
- VITE_API_BASE_URL：API 基础地址（可缺省；缺省时使用同源网关的 /api/v1/*）
- VITE_DEV_PROXY_TARGET：仅在直连前端 dev server（5173）时使用，用于将 `/api/*` 代理到后端（docker compose 默认 `http://backend:8080`）

MySQL（主从）：
- MYSQL_ROOT_PASSWORD
- MYSQL_DATABASE
- MYSQL_REPLICATION_USER
- MYSQL_REPLICATION_PASSWORD

注意：MySQL 容器首次初始化后会把 root 密码写入数据卷；如果后续修改了 `MYSQL_ROOT_PASSWORD`，需要先执行 `sudo docker compose -f deploy/docker-compose.yml down -v --remove-orphans` 清理数据卷后再重新初始化。

说明：数据库使用 Docker 官方 MySQL 镜像（`mysql:8.0`），主从复制由 deploy/mysql 中脚本进行初始化与配置。

通用端口（可选）：
- GATEWAY_PORT, BACKEND_PORT, FRONTEND_PORT, PROMETHEUS_PORT, GRAFANA_PORT, LOKI_PORT, MYSQL_MASTER_PORT, MYSQL_REPLICA_PORT

## 启动与部署

所有基础服务与脚本统一在仓库根目录第一级 deploy/ 下，通过 Docker 方式部署。

前置条件：
- 当前用户需要具备访问 Docker Daemon 的权限（能执行 `docker ps`）。如果没有权限：
  - 方式一：使用 sudo 执行脚本（例如 `sudo bash ./deploy/deploy.sh`）
  - 方式二：将用户加入 docker 组后重新登录（例如 `sudo usermod -aG docker $USER`）
- 如果本机端口冲突，可通过环境变量覆盖 `MYSQL_MASTER_PORT`/`MYSQL_REPLICA_PORT`（默认 33306/33307）。
- 启动脚本会在启动前尝试释放本项目需要的端口（先停掉占用端口的容器，再对占用端口的进程发送 TERM/KILL）。基础设施与应用会分别只释放各自需要的端口，避免互相影响。

### 一键启动（基础设施 → bootstrap → 应用）

```bash
bash ./deploy/deploy.sh
```

### 分步启动

1) 启动基础设施

```bash
bash ./deploy/infra-up.sh
```

2) 初始化（迁移 + 种子数据）与主从验证

```bash
bash ./deploy/bootstrap.sh
bash ./deploy/mysql/smoke-test.sh
```

说明：数据库初始化（建表/默认数据）通过 deploy 中的 SQL 执行完成，不依赖后端镜像/二进制参与初始化流程。

3) 启动应用（后端 → 前端）

```bash
bash ./deploy/app-up.sh
```

## 回滚与停止

- 停止应用：

```bash
bash ./deploy/app-down.sh
```

- 停止全部：

```bash
bash ./deploy/infra-down.sh
```

回滚策略：
- 网关/监控/日志为旁路组件，可先停用不影响核心业务
- 从库异常时后端读操作回退主库；必要时只保留主库读写
- 后端纯服务化后，回滚仅需替换后端镜像/二进制，不影响基础设施
