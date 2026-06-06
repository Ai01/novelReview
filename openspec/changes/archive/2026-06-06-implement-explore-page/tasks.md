## 1. 数据库层

- [x] 1.1 在 `deploy/mysql/schema.sql` 中新增 `books` 表（id, title, author, cover, description, category, tags, status, last_updated, created_at, updated_at）
- [x] 1.2 在 `deploy/mysql/schema.sql` 中新增 `comments` 表（id, book_id, user_id, content, likes, created_at），含外键约束和索引
- [x] 1.3 在 `deploy/mysql/schema.sql` 中插入种子数据：至少 20 本图书 + 40 条评论（使用 INSERT IGNORE 保证幂等）

## 2. 后端 API

- [x] 2.1 在 `backend/models.go` 中新增 `Book` 和 `Comment` 结构体（GORM 模型）
- [x] 2.2 在 `backend/handlers.go` 中实现 `ExploreHandler`：游标分页查询评论列表，关联 Book/User，返回统一 JSON 格式
- [x] 2.3 在 `backend/handlers.go` 中实现 `ExploreSearchHandler`：关键词搜索书籍+评论，支持游标分页
- [x] 2.4 在 `backend/main.go` 中注册 `GET /api/v1/explore` 和 `GET /api/v1/explore/search` 路由
- [x] 2.5 在 `backend/main.go` 中**移除**AutoMigrate，数据库完全由deploy/mysql/schema.sql管理

## 3. 前端页面

- [x] 3.1 创建 `frontend/src/pages/ExplorePage.jsx`：页面容器，包含搜索框、无限滚动列表、加载/错误/空状态处理
- [x] 3.2 创建 `frontend/src/components/ReviewCard.jsx`：单条评论卡片组件（封面、书名、作者、用户名、评论内容、点赞数、时间）
- [x] 3.3 创建 `frontend/src/components/BottomNav.jsx`：底部导航栏组件（探索/书库/+/书单/个人），接收 `activeTab` 属性高亮当前页
- [x] 3.4 在 `frontend/src/App.jsx` 中将 `/explore` 路由替换为 `<ExplorePage />` 组件
- [x] 3.5 实现 IntersectionObserver 无限滚动逻辑（监听 sentinel 元素，触发游标分页加载）
- [x] 3.6 实现搜索防抖逻辑（用户输入停止 300ms 后触发搜索请求）

## 4. 部署验证

- [x] 4.1 运行 `deploy/infra-up.sh` 启动基础设施（MySQL），验证新表创建成功
- [x] 4.2 运行 `deploy/app-up.sh` 构建并启动前后端，验证服务正常启动
- [x] 4.3 浏览器访问 `/explore`，验证页面渲染、搜索、无限滚动功能
- [x] 4.4 验证 Grafana 监控面板能正常采集探索页 API 的 metrics 和 logs
