package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rPniu/all/pkg/e"
	resp "github.com/rPniu/all/pkg/response"
	"github.com/rPniu/all/pkg/util"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, resp.Failure(e.TokenInvalid))
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, resp.Failure(e.TokenInvalid))
			c.Abort()
			return
		}

		c.Set("uid", claims.Uid)

		c.Next()
	}
}
