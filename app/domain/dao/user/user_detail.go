package dao

import (
	"example/connectify/app/domain/dao"
	"time"
)

type UserDetail struct {
	UserID      int       `gorm:"column:user_id;primary_key;" json:"user_id"`
	User        User      `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Birthday    time.Time `gorm:"column:birthday" json:"birthday"`
	PhotoUrl    string    `gorm:"column:photo_url" json:"photo_url"`
	dao.BaseModel
}
