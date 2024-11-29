package models

import "time"

// Admin 管理员表
type Admin struct {
	Id        int32     // 主键
	Username  string    // 用户名
	Mobile    string    // 手机号
	CreatedAt time.Time `gorm:"time"` // 创建时间
}
