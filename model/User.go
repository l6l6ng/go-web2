package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null;comment:用户名称" json:"user_name"`
	Password string `gorm:"type:varchar(20);not null;comment:密码" json:"password"`
	role     int    `gorm:"type:int;comment:角色" json:"role"`
	Article  []Article
}
