## Context

项目处于起步阶段，需要建立基础的账户体系。目前已有产品需求 ([product.md](file:///home/bhh/Desktop/novelWeb/product/product.md))，明确了需要登录页以及后端认证支持。

## Goals / Non-Goals

**Goals:**
- 提供完整的登录 UI（React）。
- 提供 JWT 认证的后端接口（Go）。
- 使用 MySQL 存储用户信息。
- 实现 Docker Compose 容器化编排。

**Non-Goals:**
- 第三方登录（OAuth2）。
- 验证码/短信验证。
- 完善的权限管理系统（RBAC），仅实现基础身份识别。

## Decisions

- **架构方案**: 前后端分离。
  - *Rationale*: 提高开发效率，适配未来移动端 App 开发需求。
- **后端框架**: Gin + GORM (Go 1.26+)。使用 `github.com/gin-contrib/cors` 处理跨域问题。
  - *Rationale*: Gin 轻量高效，GORM 是 Go 生态最成熟的 ORM，使用最新稳定版 Go 1.26 以获得更好的性能和安全性。官方 CORS 中间件比手动设置 Header 更安全可靠。
- **前端框架**: React + Vite + Tailwind CSS。
  - *Rationale*: React 生态丰富，Vite 构建快，Tailwind CSS 适合快速实现美观 UI。
- **认证方式**: JWT (JSON Web Token)。
  - *Rationale*: 无状态认证，方便水平扩展，适合前后端分离架构。
- **部署方式**: Docker Compose。
  - *Rationale*: 环境隔离，一键启动所有依赖（MySQL + Backend + Frontend）。通过配置 Docker 镜像加速器和代理解决国内网络环境下的镜像拉取问题。

## Risks / Trade-offs

- **[Risk]** 明文传输密码 → **Mitigation**: 后端使用 `bcrypt` 对密码进行哈希存储，确保数据库泄露后密码不可逆。
- **[Risk]** Token 泄露 → **Mitigation**: 设置合理的 Token 过期时间。
- **[Risk]** 数据库配置复杂 → **Mitigation**: 在 `docker-compose.yaml` 中配置初始化脚本和环境变量。
