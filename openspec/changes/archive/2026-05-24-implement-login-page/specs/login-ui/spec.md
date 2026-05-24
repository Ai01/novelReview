## ADDED Requirements

### Requirement: 登录页面 UI
前端系统 SHALL 提供一个响应式的登录页面，包含用户名和密码输入框。

#### Scenario: 渲染登录表单
- **WHEN** 用户访问 `/login` 路由
- **THEN** 页面显示包含“用户名/邮箱”输入框、“密码”输入框和“登录”按钮的表单

### Requirement: 登录状态管理
前端系统 SHALL 在登录成功后保存 Token，并跳转到主页。

#### Scenario: 登录成功跳转
- **WHEN** 用户提交登录表单并收到后端返回的有效 Token
- **THEN** 系统将 Token 存储在 `localStorage` 或 `Cookies` 中，并自动跳转至 `/explore` 页面
