package repository

import (
	dao "example/connectify/app/domain/dao/user"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserFollowingRepository interface {
	GetUserFollowing(id int) ([]dao.UserDetail, error)
	GetUserFollowers(id int) ([]dao.UserDetail, error)
	GetUserFollowingCount(id int) (int, error)
	GetUserFollowersCount(id int) (int, error)
	Follow(userFollowing *dao.UserFollowing) (dao.UserFollowing, error)
	Unfollow(id int, currentUserId int) error
	CheckHasFollowed(id int, currentUserId int) bool
}

type UserFollowingRepositoryImpl struct {
	db *gorm.DB
}

func (u UserFollowingRepositoryImpl) CheckHasFollowed(id int, currentUserId int) bool {
	var count int64 = 0
	log.Info("check has followed")
	err := u.db.Model(&dao.UserFollowing{}).Where("user_id = ? AND followed_user_id = ?", currentUserId, id).Count(&count)
	if err != nil {
		log.Error("Got and error when find userFollowing by id. Error: ", err)
		return count > 0
	}
	return count > 0
}

func (u UserFollowingRepositoryImpl) GetUserFollowing(id int) ([]dao.UserDetail, error) {
	users := []dao.UserDetail{}
	userFollowings := []dao.UserFollowing{}
	err := u.db.Model(&dao.UserFollowing{}).Where("user_id = ?", id).Find(&userFollowings).Error
	if err != nil {
		log.Error("Got and error when find userFollowings by id. Error: ", err)
		return []dao.UserDetail{}, err
	}
	for _, user := range userFollowings {
		userDetail := dao.UserDetail{}
		err := u.db.Model(&dao.UserDetail{}).Preload("User").Where("user_id = ?", user.ID).Find(&userDetail).Error
		if err != nil {
			log.Error("Got an error when counting comments. Error: ", err)
			return []dao.UserDetail{}, err
		}
		users = append(users, userDetail)
	}
	return users, nil
}

func (u UserFollowingRepositoryImpl) GetUserFollowers(id int) ([]dao.UserDetail, error) {
	users := []dao.UserDetail{}
	userFollowings := []dao.UserFollowing{}
	err := u.db.Model(&dao.UserFollowing{}).Where("followed_user_id = ?", id).Find(&userFollowings).Error
	if err != nil {
		log.Error("Got and error when find userFollowers by id. Error: ", err)
		return []dao.UserDetail{}, err
	}
	for _, user := range userFollowings {
		userDetail := dao.UserDetail{}
		err := u.db.Model(&dao.UserDetail{}).Preload("User").Where("user_id = ?", user.UserID).Find(&userDetail).Error
		if err != nil {
			log.Error("Got an error when counting comments. Error: ", err)
			return []dao.UserDetail{}, err
		}
		users = append(users, userDetail)
	}
	return users, nil
}

func (u UserFollowingRepositoryImpl) GetUserFollowingCount(id int) (int, error) {
	var count int64 = 0
	err := u.db.Model(&dao.UserFollowing{}).Where("user_id = ?", id).Count(&count).Error
	if err != nil {
		log.Error("Got and error when find userFollowing by id. Error: ", err)
		return int(count), err
	}
	return int(count), nil
}

func (u UserFollowingRepositoryImpl) GetUserFollowersCount(id int) (int, error) {
	var count int64 = 0
	err := u.db.Model(&dao.UserFollowing{}).Where("followed_user_id = ?", id).Count(&count).Error
	if err != nil {
		log.Error("Got and error when find userFollowers by id. Error: ", err)
		return int(count), err
	}
	return int(count), nil
}

func (u UserFollowingRepositoryImpl) Follow(userFollowing *dao.UserFollowing) (dao.UserFollowing, error) {
	following := dao.UserFollowing{}
	u.db.Where("user_id = ? AND followed_user_id = ?", userFollowing.UserID, userFollowing.FollowedUserID).Find(&following)
	if following.ID != 0 {
		log.Error("Already followed")
		return following, nil
	} else {
		err := u.db.Create(userFollowing).Error
		if err != nil {
			log.Error("Got an error when saving userFollowing. Error: ", err)
			return dao.UserFollowing{}, err
		}
	}
	return *userFollowing, nil
}

func (u UserFollowingRepositoryImpl) Unfollow(id int, currentUserId int) error {
	err := u.db.Unscoped().Where("user_id = ? AND followed_user_id = ?", currentUserId, id).Delete(&dao.UserFollowing{}).Error
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
