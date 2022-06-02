package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;comment:标签名称" json:"name"`
	Article []Article
}
