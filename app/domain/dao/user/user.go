package dao

import "example/connectify/app/domain/dao"

type User struct {
	ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Username string `gorm:"column:username;unique" json:"username"`
	Email    string `gorm:"column:email;unique" json:"email"`
	Password string `gorm:"column:password;" json:"password"`
	dao.BaseModel
}
