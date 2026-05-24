## 1. 基础架构与环境搭建

- [x] 1.1 初始化 Go 后端项目结构 (go mod init)
- [x] 1.2 初始化 React 前端项目结构 (create vite)
- [x] 1.3 编写 `docker-compose.yaml` 配置 MySQL、后端和前端容器
- [x] 1.4 配置后端 GORM 连接 MySQL 的基础代码

## 2. 后端开发 (Go)

- [x] 2.1 定义 User 模型并实现数据库自动迁移 (AutoMigrate)
- [x] 2.2 实现 JWT 工具函数 (生成与验证 Token)
- [x] 2.3 编写登录接口逻辑 (密码校验与 Token 返回)
- [x] 2.4 配置 Gin 路由并添加基础中间件 (如 CORS)

## 3. 前端开发 (React)

- [x] 3.1 安装 axios, react-router-dom, tailwindcss 等依赖
- [x] 3.2 创建登录页面 UI 组件 (Form, Inputs, Button)
- [x] 3.3 实现 API 调用逻辑与登录状态管理 (Token 存储)
- [x] 3.4 配置基础路由跳转 (登录成功后跳转至 /explore)

## 4. 联调与验证

- [x] 4.1 运行 Docker Compose 启动全栈服务
- [x] 4.2 手动测试登录流程 (前端 -> 后端 -> 数据库)
- [x] 4.3 修复联调中发现的问题 (包括 Go 版本升级、Docker 代理配置、npm 超时处理)
