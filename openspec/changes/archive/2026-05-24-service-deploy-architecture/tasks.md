## 1. Baseline & Audit

- [x] 1.1 梳理后端启动流程：定位“准备逻辑”（准备数据、初始化/迁移、数据库启停等）所在位置与调用路径
- [x] 1.2 梳理数据库访问点：标记读/写路径（Repository/DAO/Service 层）并列出需要读写分离改造的范围
- [x] 1.3 确认当前配置方式（env/config 文件）与可观测性现状（是否已有 metrics、日志输出方式）

## 2. Deploy 目录与脚本入口

- [x] 2.1 在仓库根目录新增 deploy 目录结构（第一级目录，包含 docker compose 编排、配置、初始化脚本、统一入口脚本）
- [x] 2.2 增加 deploy 基础脚本：使用 docker compose 启动/停止基础设施、执行 bootstrap（要求幂等、可重复执行、失败返回非 0）
- [x] 2.3 提供 deploy 环境变量约定（主从 DSN、端口、网关域名等）并确保脚本可在不同环境复用

## 3. MySQL 主从（主写从读）

- [x] 3.1 在 deploy 的 docker compose 编排中增加 MySQL master 服务（持久化卷、健康检查、初始化用户/权限）
- [x] 3.2 在 deploy 的 docker compose 编排中增加 MySQL replica 服务并配置 replication（复制用户、replica 启动脚本/配置）
- [x] 3.3 增加主从验证的 smoke test（写入主库后在从库可见；记录可接受的延迟与失败行为）

## 4. Nginx 网关

- [x] 4.1 在 deploy 中增加 Nginx 配置与容器编排：/api 反代后端、非 /api 路由服务前端静态资源
- [x] 4.2 增加网关健康检查端点，并确保网关可在后端未启动时仍可健康响应
- [x] 4.3 配置并验证标准代理头透传（X-Forwarded-For、X-Forwarded-Proto 等）

## 5. Prometheus + Grafana + Loki（监控/日志）

- [x] 5.1 确认/实现后端 metrics 暴露端点（Prometheus text format）
- [x] 5.2 在 deploy 中增加 Prometheus 配置并抓取后端 metrics（目标可用状态为 up）
- [x] 5.3 在 deploy 中增加 Grafana provisioning（Prometheus datasource）并提供最小可用健康面板
- [x] 5.4 在 deploy 中增加 Loki 服务并确保可独立启动（不依赖后端/前端）
- [x] 5.5 在 deploy 中增加日志采集组件并将后端/网关等日志写入 Loki（至少支持按 service/env 标签检索）
- [x] 5.6 在 deploy 中为 Grafana 增加 Loki datasource，并验证日志可查询（按服务筛选、关键字检索）

## 6. 后端纯服务化（剥离准备逻辑）

- [x] 6.1 将数据库初始化/迁移/种子数据等准备逻辑从后端启动流程中移除并迁移到 deploy/bootstrap 脚本
- [x] 6.2 增加后端 readiness 信号：至少基于写库连通性；写库不可用时 readiness 为 not ready
- [x] 6.3 清理后端中任何“启动/管理基础设施”的行为（例如启动数据库容器/进程、自动导入准备数据）

## 7. 应用侧读写分离改造

- [x] 7.1 增加后端配置：WRITE_DSN（主库）与 READ_DSN（从库），并支持 READ_DSN 不可用时回退到 WRITE_DSN 读
- [x] 7.2 实现数据库路由封装（写强制走主库；读默认走从库；提供“强一致读走主库”的显式能力）
- [x] 7.3 将关键路径读请求接入路由封装并完成回归验证（写后读一致性场景覆盖）

## 8. 部署流程与文档

- [x] 8.1 在 deploy 中固化部署顺序：基础设施 → 后端 → 前端，并提供一键入口脚本
- [x] 8.2 增加整体 README：架构说明、目录结构、端口/入口、启动与部署步骤、回滚策略
