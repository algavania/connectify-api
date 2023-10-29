package dao

import (
	"example/connectify/app/domain/dao"
	user "example/connectify/app/domain/dao/user"
)

type ChatParticipant struct {
	ID     int       `gorm:"column:id; primary_key; not null" json:"id"`
	ChatID int       `gorm:"column:chat_id;" json:"chat_id"`
	Chat   Chat      `gorm:"foreignkey:ChatID;constraint:OnDelete:CASCADE;" json:"-"`
	UserID int       `gorm:"column:user_id;" json:"user_id"`
	User   user.User `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	dao.BaseModel
}
