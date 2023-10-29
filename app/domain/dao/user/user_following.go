package dao

import "example/connectify/app/domain/dao"

type UserFollowing struct {
	ID             int  `gorm:"column:id; primary_key; not null" json:"id"`
	UserID         int  `gorm:"column:user_id;" json:"user_id"`
	User           User `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	FollowedUserID int  `gorm:"column:followed_user_id;" json:"followed_user_id"`
	FollowedUser   User `gorm:"foreignkey:FollowedUserID;constraint:OnDelete:CASCADE;" json:"-"`
	dao.BaseModel
}
