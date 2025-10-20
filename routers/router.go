package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rPniu/all/middleware"
	"github.com/rPniu/all/pkg/setting"
	"github.com/rPniu/all/routers/api"
	v1 "github.com/rPniu/all/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RateLimitMiddleware())
	gin.SetMode(setting.RunMode)

	//注册新用户
	r.POST("Register", api.Register)
	//登录
	r.POST("Login", api.Login)

	r.Static("/static", "./static")

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.GET("/test", v1.Test)
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "***404 not found***")
	})

	return r
}
