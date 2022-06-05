package model

import (
	"github.com/l6l6ng/go-web2/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	UserName string `gorm:"type:varchar(20);not null;comment:用户名称" json:"user_name" form:"user_name" binding:"required"`
	Password string `gorm:"type:varchar(20);not null;comment:密码" json:"password" form:"password" binding:"required"`
	role     int    `gorm:"type:int;comment:角色" json:"role" form:"role" binding:"required"`
	Article  []Article
	CommonTimestampsField
	CommonSoftDeletesField
}

// UserExit 用户是否存在
func UserExit(userName string) int {
	user := User{}
	db.Select("id").Where("user_name = ?", userName).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// AddUser 创建用户
func AddUser(user *User) int {
	err := db.Create(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}
