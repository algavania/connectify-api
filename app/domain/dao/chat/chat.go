package dao

import "example/connectify/app/domain/dao"

type Chat struct {
	ID   int    `gorm:"column:id; primary_key; not null" json:"id"`
	Type int    `gorm:"column:type;" json:"type"`
	Name string `gorm:"column:name;" json:"name"`
	dao.BaseModel
}
