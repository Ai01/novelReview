## 1. 创建目录结构

- [x] 1.1 创建 `internal/` 下 6 个子目录

## 2. 迁移 Domain 层（按实体拆分子包）

- [x] 2.1 `domain/user/user.go` — User 结构体
- [x] 2.2 `domain/book/book.go` — Book 结构体
- [x] 2.3 `domain/comment/comment.go` — Comment 结构体（import user + book）

## 3. 迁移基础设施层

- [x] 3.1 创建 `internal/infra/database.go`

## 4. 迁移 Repository 层

- [x] 4.1 `repository/user_repo.go` — UserRepository 接口+实现
- [x] 4.2 `repository/comment_repo.go` — CommentRepository 接口+实现

## 5. 迁移 Service 层

- [x] 5.1 `service/auth_service.go` — AuthService（JWT+登录逻辑）
- [x] 5.2 `service/explore_service.go` — ExploreService（探索页逻辑）

## 6. 迁移 Handler 层（按实体拆分）

- [x] 6.1 `handler/auth_handler.go` — LoginHandler
- [x] 6.2 `handler/explore_handler.go` — ExploreHandler + ExploreSearchHandler

## 7. 创建 Router 层

- [x] 7.1 `router/router.go` — 集中注册所有路由

## 8. 重构 main.go

- [x] 8.1 重写 `main.go`，组装入口

## 9. 清理旧文件

- [x] 9.1 删除旧文件

## 10. 验证编译和部署

- [x] 10.1 `go build` 编译零错误
- [x] 10.2 部署验证 API 行为不变

## 11. 更新 README.md

- [x] 11.1 新增后端架构章节
- [x] 11.2 更新目录结构描述
