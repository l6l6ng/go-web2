package model

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null;comment:标签名称" json:"name"`
	Article []Article
	CommonTimestampsField
	CommonSoftDeletesField
}
