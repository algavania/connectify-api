package dao

import (
	"example/connectify/app/domain/dao"
	user "example/connectify/app/domain/dao/user"
)

type Post struct {
	ID           int       `gorm:"column:id; primary_key; not null" json:"id"`
	UserID       int       `gorm:"column:user_id;" json:"user_id"`
	User         user.User `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	ParentPostID int       `gorm:"column:parent_post_id;" json:"parent_post_id"`
	ParentPost   user.User `gorm:"foreignkey:ParentPostID;constraint:OnDelete:CASCADE;" json:"-"`
	Content      string    `gorm:"column:content;" json:"content"`
	Media1       string    `gorm:"column:media1;" json:"media1"`
	Media2       string    `gorm:"column:media2;" json:"media2"`
	Media3       string    `gorm:"column:media3;" json:"media3"`
	Media4       string    `gorm:"column:media4;" json:"media4"`
	dao.BaseModel
}
