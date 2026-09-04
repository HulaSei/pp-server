# 共享包与应用实现的归属

`pkg` 保留 12 个一级目录：`cache`、`conf`、`httpx`、`logger`、`orm`、`random`、
`requestmeta`、`slicesx`、`templatex`、`timeutil`、`trace`、`xerr`。
这些包不得反向依赖应用或业务模块；`internal/arch` 检查这一边界。

## 实际合并

- HTTP 参数绑定与响应包装统一到 `pkg/httpx`，保留原有 JSON 与 Swagger schema 名称。
- AES 加解密内聚到 `internal/auth/deviceauth`，签名与密文格式不变。
- 邮件与唯一的 SMTP 实现统一到 `internal/mail`。
- 服务组、停机钩子统一到 `internal/lifecycle`，保留停止顺序、只执行一次及 panic 恢复行为。
- trace ID/Span ID 直接使用 OpenTelemetry SDK，移除重复的上下文包，避免日志循环依赖。
- 邮箱、手机号与标识规范化统一到 `internal/auth/identifier`。
- Telegram webhook 密钥推导并入 notification 模块。

## 应用实现下沉

- 支付协议和汇率适配器归 billing；应用组装与刷新任务通过 billing 的公开入口使用汇率缓存。
- OAuth 供应商和 state 管理归 identity；不同供应商仍使用独立子包。
- IP 定位与节点倍率归 network；组装层通过 network 的公开入口持有倍率管理器。
- 队列追踪适配器、WebSocket 设备管理、邮件和短信放在共享的应用基础设施层。
- JWT、设备会话、验证码挑战和限流保留独立的认证子边界，不合成一个万能认证包。
- 应用版本与上下文键放在 `internal/constant`，构建脚本的版本注入路径同步更新。

## 原工具集合的拆分

密码哈希归认证；字段映射放在 `internal/mapping`；系统配置转换放在 `internal/config`；
协议签名和兼容编码放在 `internal/protocolkey`；订单号生成归订单实体；金额格式化归支付。
通用切片、日期和 ETag 分别并入 `slicesx`、`timeutil`、`httpx`。

本轮调整不改变数据库结构、HTTP 接口、支付签名、设备加密、邀请码或订阅 Token 算法。
供应商适配器及安全子包的移动并不等于删掉了它们；统计目录数量时，应区分公共包收拢
和实际 Go 包合并。外部 Go 项目若直接导入过被迁移的服务端包，需要调整，不能导入
服务端的 `internal` 实现；HTTP 客户端协议不受这次目录调整影响。
