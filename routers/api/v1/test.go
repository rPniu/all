package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rPniu/all/models"
	"github.com/rPniu/all/pkg/e"
	"github.com/rPniu/all/pkg/logging"
	resp "github.com/rPniu/all/pkg/response"
	"net/http"
)

func Test(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		logging.Info("No UID found in context")
		c.JSON(http.StatusUnauthorized, resp.Failure(e.TokenInvalid))
		return
	}

	uidStr, ok := uid.(string)
	if !ok {
		logging.Error("UID is not a string type")
		c.JSON(http.StatusOK, resp.Failure(e.UserNotExist))
		return
	}

	// 使用新的GetUserInfo函数获取用户信息
	userInfo, err := models.GetUserInfo(uidStr)
	if err != nil {
		logging.Errorf("Failed to get user info: %v", err)
		c.JSON(http.StatusOK, resp.Failure(e.UserNotExist))
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, resp.Success(userInfo))
}
