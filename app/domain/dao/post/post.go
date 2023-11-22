package dao

import (
	"example/connectify/app/domain/dao"
	user "example/connectify/app/domain/dao/user"
)

type Post struct {
	ID           int       `gorm:"column:id; primary_key; not null" json:"id"`
	UserID       int       `gorm:"column:user_id;" json:"user_id"`
	User         user.User `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	ParentPostID *int      `gorm:"column:parent_post_id;default:null" json:"parent_post_id"`
	ParentPost   []Post    `gorm:"foreignkey:ParentPostID;constraint:OnDelete:CASCADE;" json:"-"`
	Content      string    `gorm:"column:content;" json:"content"`
	Media1       *string   `gorm:"column:media1;default:null" json:"media1"`
	Media2       *string   `gorm:"column:media2;default:null" json:"media2"`
	Media3       *string   `gorm:"column:media3;default:null" json:"media3"`
	Media4       *string   `gorm:"column:media4;default:null" json:"media4"`
	dao.BaseModel
}

type PostResponse struct {
	Post
	CommentCount int64 `json:"comment_count"`
}
