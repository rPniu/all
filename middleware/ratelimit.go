package middleware

import (
	"time"

	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/rPniu/all/pkg/e"
	resp "github.com/rPniu/all/pkg/response"
)

// RateLimitMiddleware 创建一个用于请求限流的 Gin 中间件。
// 它使用内存存储来跟踪每个 IP 地址的请求计数。
func RateLimitMiddleware() gin.HandlerFunc {
	// 为限流器创建一个内存存储。
	// 此配置允许每秒 5 个请求。
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second * 3, // 令牌桶中添加令牌的速率（例如，每 3 秒 1 个令牌）
		Limit: 1,               // 令牌桶可以容纳的最大令牌数（例如，1 个请求的突发容量）
	})

	// 使用自定义错误处理和密钥提取创建限流中间件。
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		// ErrorHandler 在请求被限流时调用。
		ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
			c.JSON(429, resp.FailureAddDetail(e.TooManyRequests, "每3s可请求一次"))
		},
		// KeyFunc 用于提取客户端标识符进行限流。
		// 在此示例中，它使用客户端的 IP 地址。
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	})

	return mw
}
