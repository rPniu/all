package e

// 通用错误 (1000 ~ 1099)
const (
	Success      = 1000 // 成功
	InvalidParam = 1001 // 参数错误
	DBError      = 1002 // 数据库错误
	UnknownError = 1003 // 未知错误
	ServiceDown  = 1004 // 服务不可用
)

// 用户相关 (2000 ~ 2099)
const (
	UserNotLogin  = 2000 // 用户未登录
	UserNotExist  = 2001 // 用户不存在
	PasswordError = 2002 // 密码错误
	UserDisabled  = 2003 // 用户被禁用
	TokenInvalid  = 2004 // Token 无效或过期
)

// 权限相关 (3000 ~ 3099)
const (
	NoPermission   = 3000 // 无权限访问
	RoleNotAllowed = 3001 // 角色权限不足
	AuthFailed     = 3002 // 认证失败
)

// 资源相关 (4000 ~ 4099)
const (
	ResourceNotFound   = 4000 // 资源不存在
	ResourceExists     = 4001 // 资源已存在
	UploadFailed       = 4002 // 文件上传失败
	InvalidImageFormat = 4003 // 图片格式错误
	FileTooLarge       = 4004 // 文件大小超限
)

// 限流相关 (6000 ~ 6099)
const (
	TooManyRequests     = 6000 // 请求过多，触发限流
	LoginTooFrequent    = 6001 // 登录请求过于频繁
	RegisterTooFrequent = 6002 // 注册请求过于频繁
	IPBlocked           = 6003 // IP 被临时封禁
	AccountLocked       = 6004 // 账号因频繁操作被锁定
	APIQuotaExceeded    = 6005 // API 调用配额超限
)
