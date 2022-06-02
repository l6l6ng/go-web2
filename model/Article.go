package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null;comment:标题" json:"title"`
	CategoryId int    `gorm:"type;int;not null;column:cid;comment:标签编号" json:"c_id""`
	Desc       string `gorm:"type:varchar(200);comment:描述" json:"desc"`
	Content    string `gorm:"type:text;comment:内容" json:"content"`
	Img        string `gorm:"type:varchar(100);comment:图片" json:"img"`
	UserId     int    `gorm:"type:int;not null;comment:用户编号" json:"user_id"`
}
