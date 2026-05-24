## ADDED Requirements

### Requirement: 用户登录接口
后端系统 SHALL 提供一个 RESTful API 接口，用于验证用户凭据并返回 JWT。

#### Scenario: 登录成功
- **WHEN** 用户提供正确的用户名和密码发送 POST 请求到 `/api/v1/login`
- **THEN** 系统返回 HTTP 200 状态码，并附带包含 Token 的 JSON 响应

#### Scenario: 登录失败（密码错误）
- **WHEN** 用户提供错误的密码发送 POST 请求到 `/api/v1/login`
- **THEN** 系统返回 HTTP 401 状态码，并提示“用户名或密码错误”

### Requirement: 用户数据持久化
系统 SHALL 使用 MySQL 存储用户信息，密码必须进行加密存储。系统启动时 SHALL 自动初始化默认管理员账户。

#### Scenario: 存储加密密码
- **WHEN** 用户注册或修改密码时
- **THEN** 数据库中存储的 `password` 字段必须是经过 `bcrypt` 哈希后的字符串，而非明文

#### Scenario: 自动初始化管理员
- **WHEN** 系统首次启动且数据库中不存在 admin 用户时
- **THEN** 系统 SHALL 自动创建用户名 admin，密码 admin123 的初始账户
