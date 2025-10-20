package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rPniu/all/models"
	"github.com/rPniu/all/pkg/e"
	"github.com/rPniu/all/pkg/logging"
	"github.com/rPniu/all/pkg/response"
	"github.com/rPniu/all/pkg/util"
	"net/http"
)

func Register(c *gin.Context) {
	var code int

	password := c.PostForm("password")

	newUser, err := models.CreateUser(password)
	if err != nil {
		code = e.DBError
		logging.Errorf("create user failed: %v", err)
		c.JSON(http.StatusOK, resp.Failure(code))
		return
	} else {
		c.JSON(http.StatusOK, resp.Success(newUser))
	}
}

func Login(c *gin.Context) {
	uid := c.PostForm("uid")
	password := c.PostForm("password")

	if uid == "" || password == "" { // 补充：参数合法性校验
		logging.Warnf("login param missing: uid=%s, password=%s", uid, password)
		c.JSON(http.StatusOK, resp.Failure(e.InvalidParam))
		return
	}

	user, ok := models.CheckAuth(uid, password)
	if ok {
		token, _ := c.Cookie("token")
		if token != "" {
			_, err := util.ParseToken(token)
			if err == nil {
				c.JSON(http.StatusOK, resp.Success(user))
				return
			}
		}
		newToken, err := util.GenerateToken(uid)
		if err != nil {
			logging.Errorf("generate token failed: %v, uid: %s", err, uid)
			c.JSON(http.StatusOK, resp.Failure(e.UnknownError))
			return
		}
		c.SetCookie("token",
			newToken,
			3600,
			"/",
			"localhost",
			false,
			true)
		logging.Infof("login success, uid: %s", uid)
		c.JSON(http.StatusOK, resp.Success(user))
	} else {
		if user != nil {
			logging.Infof("PasswordError, uid: %s", uid)
			c.JSON(http.StatusOK, resp.Failure(e.PasswordError))
		} else {
			logging.Infof("UserNotExist, uid: %s", uid)
			c.JSON(http.StatusOK, resp.Failure(e.UserNotExist))
		}
	}
}
