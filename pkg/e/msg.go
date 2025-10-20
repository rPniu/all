package e

var MsgFlags = map[int]string{
	Success:      "成功",
	InvalidParam: "参数错误",
	DBError:      "数据库错误",
	UnknownError: "未知错误",
	ServiceDown:  "服务不可用",

	UserNotLogin:  "用户未登录",
	UserNotExist:  "用户不存在",
	PasswordError: "密码错误",
	UserDisabled:  "用户被禁用",
	TokenInvalid:  "Token 无效或过期",

	NoPermission:   "无权限访问",
	RoleNotAllowed: "角色权限不足",
	AuthFailed:     "认证失败",

	ResourceNotFound:   "资源不存在",
	ResourceExists:     "资源已存在",
	UploadFailed:       "文件上传失败",
	InvalidImageFormat: "图片格式错误",
	FileTooLarge:       "文件大小超限",

	TooManyRequests:     "请求过多，触发限流",
	LoginTooFrequent:    "登录请求过于频繁",
	RegisterTooFrequent: "注册请求过于频繁",
	IPBlocked:           "IP 被临时封禁",
	AccountLocked:       "账号因频繁操作被锁定",
	APIQuotaExceeded:    "API 调用配额超限",
}

// GetMsg 获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[UnknownError]
}
