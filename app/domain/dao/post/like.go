package dao

import (
	"example/connectify/app/domain/dao"
	user "example/connectify/app/domain/dao/user"
)

type Like struct {
	ID     int       `gorm:"column:id; primary_key; not null" json:"id"`
	UserID int       `gorm:"column:user_id;" json:"user_id"`
	User   user.User `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	PostID int       `gorm:"column:post_id;" json:"post_id"`
	Post   Post      `gorm:"foreignkey:PostID;constraint:OnDelete:CASCADE;" json:"-"`
	dao.BaseModel
}
