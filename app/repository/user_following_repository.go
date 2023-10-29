package repository

import (
	dao "example/connectify/app/domain/dao/user"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserFollowingRepository interface {
	GetUserFollowing(id int) (dao.UserFollowing, error)
	GetUserFollowers(id int) (dao.UserFollowing, error)
	Follow(userFollowing *dao.UserFollowing) (dao.UserFollowing, error)
	Unfollow(userFollowing *dao.UserFollowing) error
}

type UserFollowingRepositoryImpl struct {
	db *gorm.DB
}

func (u UserFollowingRepositoryImpl) GetUserFollowing(id int) (dao.UserFollowing, error) {
	userFollowing := dao.UserFollowing{
		UserID: id,
	}
	err := u.db.Find(&userFollowing).Error
	if err != nil {
		log.Error("Got and error when find userFollowing by id. Error: ", err)
		return dao.UserFollowing{}, err
	}
	return userFollowing, nil
}

func (u UserFollowingRepositoryImpl) GetUserFollowers(id int) (dao.UserFollowing, error) {
	userFollowing := dao.UserFollowing{
		FollowedUserID: id,
	}
	err := u.db.Find(&userFollowing).Error
	if err != nil {
		log.Error("Got and error when find userFollowers by id. Error: ", err)
		return dao.UserFollowing{}, err
	}
	return userFollowing, nil
}

func (u UserFollowingRepositoryImpl) Follow(userFollowing *dao.UserFollowing) (dao.UserFollowing, error) {

	data, err := u.GetUserFollowing(userFollowing.ID)
	if err != nil {
		err = u.db.Create(userFollowing).Error
	} else {
		userFollowing.CreatedAt = data.CreatedAt
		err = u.db.Updates(userFollowing).Error
	}
	if err != nil {
		log.Error("Got an error when saving userFollowing. Error: ", err)
		return dao.UserFollowing{}, err
	}
	return *userFollowing, nil
}

func (u UserFollowingRepositoryImpl) Unfollow(userFollowing *dao.UserFollowing) error {
	err := u.db.Unscoped().Delete(&dao.UserFollowing{}, userFollowing).Error
	if err != nil {
		log.Error("Got an error when delete userFollowing. Error: ", err)
		return err
	}
	return nil
}

func UserFollowingRepositoryInit(db *gorm.DB) *UserFollowingRepositoryImpl {
	db.AutoMigrate(&dao.UserFollowing{})
	return &UserFollowingRepositoryImpl{
		db: db,
	}
}
