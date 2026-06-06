## Why

探索页是用户登录后的默认首页和内容发现入口，当前 `/explore` 路由仅为占位文字。需要实现完整的探索页以支撑产品核心的书籍发现与点评浏览体验，衔接已有的登录体系和部署架构。

## What Changes

- **前端**：实现探索页完整 UI，包含顶部搜索框、无限滚动书籍点评信息流、全局底部导航栏（探索/书库/+/书单/个人）
- **后端**：新增探索页 API（书籍点评列表查询、关键词搜索），支持分页/游标分页
- **数据库**：新增 `books` 表和 `comments` 表，补充种子数据
- **部署**：无需新增基础设施，在现有 Docker Compose 架构上增量部署

## Capabilities

### New Capabilities
- `explore-ui`: 探索页前端界面，包含搜索框、无限滚动信息流、底部导航栏，支持移动端和 PC 端
- `explore-api`: 探索页后端 API，提供书籍点评列表的分页查询和关键词搜索接口
- `explore-db`: 探索页相关的数据库表结构（books、comments）及种子数据
- `explore-deploy`: 探索页功能相关的部署配置更新

### Modified Capabilities
<!-- 无现有 spec 需要修改 -->

## Impact

- **前端**：`frontend/src/App.jsx` 路由替换，新增 `ExplorePage.jsx` 及相关组件
- **后端**：`backend/handlers.go` 新增探索页 handler，`backend/models.go` 新增 Book/Comment 模型，`backend/main.go` 注册新路由
- **数据库**：`deploy/mysql/schema.sql` 新增建表语句和种子数据
- **部署**：无架构变更，标准 `deploy.sh` 流程即可上线
