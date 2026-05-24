## Why

本项目旨在构建一个小说评论社区，登录功能是系统的核心入口，用于识别用户身份、保护用户隐私以及支持个性化功能（如书评发布、收藏书籍/书单）。

## What Changes

- 实现前后端分离的登录功能。
- **前端 (React)**：创建登录页面，支持输入邮箱/用户名和密码，处理登录状态及 Token 存储。
- **后端 (Go + Gin)**：提供登录接口，集成 MySQL 数据库，使用 GORM 进行数据操作，支持 JWT 认证。
- **部署 (Docker)**：配置 Dockerfile 和 Docker Compose，一键启动 MySQL、后端和前端服务。

## Capabilities

### New Capabilities
- `user-auth`: 用户认证系统，包含用户注册、登录、Token 生成与验证。
- `login-ui`: 登录界面，包含表单验证、错误提示及跳转逻辑。

### Modified Capabilities
<!-- 无 -->

## Impact

- **API**: 新增 `/api/v1/login` 接口。
- **Database**: 新增 `users` 表。
- **DevOps**: 引入 Docker Compose 编排，影响本地开发与部署流程。
- **Dependencies**: 引入 `github.com/gin-gonic/gin`, `gorm.io/gorm`, `golang-jwt/jwt` (Go) 和 `axios`, `react-router-dom` (React)。
