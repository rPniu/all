package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rPniu/all/pkg/logging"
	"github.com/rPniu/all/pkg/util"
)

// CheckAuth 验证用户身份并返回用户信息
func CheckAuth(userID, password string) (*UserInfo, bool) {
	id, err := util.StrToInt(userID)
	if err != nil {
		logging.Info(fmt.Sprintf("id %s 转换失败", userID))
		return nil, false
	}

	var user UserCore
	result := db.First(&user, id)

	fmt.Println(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logging.Infof("用户ID=%d不存在", id)
			return nil, false
		}
		logging.Infof("查询用户ID=%v失败: %v", id, result.Error)
		return nil, false
	}

	// 检查密码哈希是否为空
	if user.PasswordHash == "" {
		logging.Infof("用户ID=%d的密码哈希为空", id)
		return nil, false
	}

	// 安全地比较密码
	isValid := util.ComparePasswordHashAndPassword(user.PasswordHash, password)

	// 只返回非敏感的用户信息
	if isValid {
		return user.ToUserInfo(), true
	}

	return nil, false
}

// CreateUser 创建新用户并返回用户信息
func CreateUser(password string) (*UserInfo, error) {
	passwordHash, err := util.GetPasswordHash(password)
	if err != nil {
		return nil, err
	}
	user := &UserCore{
		UserAccount: UserAccount{
			PasswordHash: passwordHash,
		},
	}
	result := db.Create(&user)
	if result.Error != nil {
		logging.Error(result.Error.Error())
		return nil, result.Error
	}

	// 返回非敏感的用户信息
	return user.ToUserInfo(), nil
}
