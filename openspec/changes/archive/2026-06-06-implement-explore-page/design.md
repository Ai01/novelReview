## Context

探索页当前为占位页面（`<h1>探索页 (开发中)</h1>`），需要实现为完整功能页面。后端已有 Go/Gin/GORM 架构、JWT 鉴权体系、MySQL 主从读写分离。前端已有 React + Vite + TailwindCSS + React Router 基础框架。本设计基于现有架构增量开发。

## Goals / Non-Goals

**Goals:**
- 实现探索页完整 UI：搜索框、无限滚动书籍点评信息流、底部导航栏
- 实现后端 API：书籍点评列表查询（游标分页）、关键词搜索
- 创建 books 和 comments 数据库表及种子数据
- 前后端联调通过，支持移动端和 PC 端响应式布局

**Non-Goals:**
- 不实现书籍详情页（后续变更）
- 不实现书库页、书单页、个人中心页
- 不实现底部 + 号弹窗
- 不实现点评发布功能（探索页只读）

## Decisions

### 1. 分页策略：游标分页（Cursor-based）

**选择**：使用 `comment.id` 作为游标，`GET /api/v1/explore?cursor=0&limit=20`。

**理由**：无限滚动场景下游标分页比 offset 分页更适合——数据插入不会导致重复或遗漏。

**替代方案**：Offset 分页更简单，但在高频写入时会出现数据漂移。探索页作为内容发现入口，数据一致性要求较低，但游标分页体验更好且实现成本不高。

### 2. API 端点设计

```
GET  /api/v1/explore          # 探索信息流（默认按评论时间倒序）
GET  /api/v1/explore/search   # 关键词搜索书籍+评论
```

两个端点均不需要登录鉴权（探索页可公开访问）。

**响应格式：**
```json
{
  "items": [
    {
      "comment_id": 1,
      "book": { "id": 1, "title": "书名", "author": "作者", "cover": "url" },
      "user": { "id": 1, "username": "用户", "avatar": "url" },
      "content": "评论内容",
      "likes": 10,
      "created_at": "2026-05-23T12:00:00Z"
    }
  ],
  "next_cursor": 20,
  "has_more": true
}
```

### 3. 数据模型新增

**Book 表：**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 自增主键 |
| title | VARCHAR(255) | 书名 |
| author | VARCHAR(255) | 作者 |
| cover | VARCHAR(512) | 封面图 URL |
| description | TEXT | 简介 |
| category | VARCHAR(100) | 分类 |
| tags | VARCHAR(255) | 标签（逗号分隔） |
| status | VARCHAR(50) | 连载中/已完结 |
| last_updated | DATETIME | 最后更新时间 |
| created_at | DATETIME | 创建时间 |

**Comment 表：**
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 自增主键（作为分页游标） |
| book_id | BIGINT FK | 关联书籍 |
| user_id | BIGINT FK | 评论用户 |
| content | TEXT | 评论内容 |
| likes | INT DEFAULT 0 | 点赞数 |
| created_at | DATETIME | 评论时间 |

**索引：**
- `comment (book_id, created_at)` — 按书籍查评论
- `comment (created_at)` — 全局时间线排序
- `book FULLTEXT(title, author)` — 全文搜索

### 4. 前端架构

- **路由**：`/explore` 替换现有占位 Route
- **组件拆分**：
  - `ExplorePage.jsx` — 页面容器（搜索框 + 列表 + 底部导航）
  - `BottomNav.jsx` — 全局底部导航栏组件（可复用于其他页面）
  - `ReviewCard.jsx` — 单条评论卡片
- **无限滚动**：IntersectionObserver 监听底部 sentinel 元素，触发加载下一页
- **状态管理**：`useState` + `useEffect`，不引入额外状态库（保持项目轻量）

### 5. 响应式设计

- 移动端（<768px）：单列布局，底部导航栏固定
- PC 端（>=768px）：内容区最大宽度 640px 居中，底部导航栏替换为底部固定栏

### 6. 搜索实现

MySQL `LIKE '%keyword%'` 配合 FULLTEXT 索引。搜索范围：书名、作者、评论内容。简单实现，后续可升级为 Elasticsearch。

## Risks / Trade-offs

- **[风险] FULLTEXT 索引对中文支持有限** → 使用 `LIKE` 兜底，当前数据量下性能可接受。后续可接入 Elasticsearch 或使用 MySQL ngram parser。
- **[风险] 种子数据不足导致页面空旷** → 种子数据至少插入 30 条评论覆盖多页滚动测试。
- **[风险] 游标分页无法跳页** → 探索页无跳页需求，纯无限滚动场景游标是最佳选择。
- **[权衡] 评论数据未做缓存** → 当前阶段直接查库，QPS 不高。后续可加 Redis 缓存热门评论。

## Open Questions

- 书籍封面图是否需要内置 CDN/图床？建议使用外部 URL 占位，后期对接对象存储。
