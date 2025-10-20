package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/rPniu/all/pkg/logging"
	"github.com/rPniu/all/pkg/util"
	"time"
)

// UserAccount 用户账号信息，包含敏感数据
type UserAccount struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;comment:用户账号ID，自增整型" json:"id,omitempty"`
	PasswordHash string    `gorm:"size:100;not null;comment:密码（加密存储）" json:"-"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

// UserCore 用户核心信息，包含账号和个人资料
type UserCore struct {
	UserAccount
	Uid         string `gorm:"-:all" json:"uid,omitempty"`
	Nickname    string `gorm:"size:50;comment:用户名" json:"nickname,omitempty"`
	Status      uint8  `gorm:"default:1;comment:状态：1正常/0禁用" json:"status,omitempty"`
	Description string `gorm:"size:255;comment:用户简介" json:"description,omitempty"`
	AvatarURL   string `gorm:"size:255;default:'/static/avatar0001.png';comment:头像URL" json:"avatar_url,omitempty"`
}

// UserInfo 用户信息，只包含需要返回给客户端的非敏感信息
type UserInfo struct {
	Uid         string    `json:"uid"`
	Nickname    string    `json:"nickname"`
	Status      uint8     `json:"status"`
	Description string    `json:"description"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func (*UserAccount) TableName() string {
	return "user_info"
}

// ToUserInfo 将UserCore转换为UserInfo
func (u *UserCore) ToUserInfo() *UserInfo {
	return &UserInfo{
		Uid:         u.Uid,
		Nickname:    u.Nickname,
		Status:      u.Status,
		Description: u.Description,
		AvatarURL:   u.AvatarURL,
		CreatedAt:   u.CreatedAt,
	}
}

func (u *UserCore) AfterFind(tx *gorm.DB) error {
	// 1. 先检查 u.ID 是否有效（避免传入 0 到 IntTo10DigitStr）
	if u.ID == 0 {
		logging.Errorf("AfterFind: 用户 ID 为 0，无法生成 Uid")
		return fmt.Errorf("invalid user ID: %d", u.ID) // 返回错误，让 GORM 捕获
	}
	// 2. 不忽略错误，打印异常
	uid, err := util.IntTo10DigitStr(u.ID)
	if err != nil {
		logging.Errorf("AfterFind: 生成 Uid 失败，ID=%d, err=%v", u.ID, err)
		return err // 返回错误，避免后续逻辑使用异常 Uid
	}
	u.Uid = uid
	return nil
}

func (u *UserCore) AfterCreate(tx *gorm.DB) error {
	if u.ID == 0 {
		logging.Errorf("AfterCreate: 用户 ID 为 0，无法生成 Uid")
		return fmt.Errorf("invalid user ID: %d", u.ID)
	}
	uid, err := util.IntTo10DigitStr(u.ID)
	if err != nil {
		logging.Errorf("AfterCreate: 生成 Uid 失败，ID=%d, err=%v", u.ID, err)
		return err
	}
	u.Uid = uid
	return nil
}

// GetUserInfo 根据用户ID获取用户信息
func GetUserInfo(userID string) (*UserInfo, error) {
	id, err := util.StrToInt(userID)
	if err != nil {
		logging.Info(fmt.Sprintf("id %s 转换失败", userID))
		return nil, err
	}

	var user UserCore
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user.ToUserInfo(), nil
}
