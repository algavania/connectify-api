package repository

import (
	dao "example/connectify/app/domain/dao/user"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserDetailRepository interface {
	FindAllUsers() ([]dao.UserDetail, error)
	FindUserById(id int) (dao.UserDetail, error)
	Save(user *dao.UserDetail) (dao.UserDetail, error)
}

type UserDetailRepositoryImpl struct {
	db *gorm.DB
}

func (u UserDetailRepositoryImpl) FindAllUsers() ([]dao.UserDetail, error) {
	log.Info("get all users")
	users := []dao.UserDetail{}
	err := u.db.Preload("User").Find(&users).Error
	log.Info("users", users)
	if err != nil {
		log.Error("Got an error when finding all user details. Error: ", err)
		return users, err
	}
	return users, nil
}

func (u UserDetailRepositoryImpl) FindUserById(id int) (dao.UserDetail, error) {
	user := dao.UserDetail{
		UserID: id,
	}
	err := u.db.First(&user).Error
	log.Info("find user by id", id)
	if err != nil || id == 0 {
		log.Error("Got and error when find user detail by id. Error: ", err)
		return dao.UserDetail{}, err
	}
	return user, nil
}

func (u UserDetailRepositoryImpl) Save(user *dao.UserDetail) (dao.UserDetail, error) {
	userDetail := dao.UserDetail{}
	err := u.db.Where("user_id = ?", user.UserID).Find(&userDetail).Error
	log.Info("userDetail: ", userDetail)
	log.Info(err)
	if userDetail.UserID == 0 {
		err = u.db.Create(user).Error
	} else {
		user.CreatedAt = userDetail.CreatedAt
		err = u.db.Updates(user).Error
		u.db.Where("user_id = ?", user.UserID).Find(user)
	}
	if err != nil {
		log.Error("Got an error when save user detail. Error: ", err)
		return dao.UserDetail{}, err
	}
	return *user, nil
}

func UserDetailRepositoryInit(db *gorm.DB) *UserDetailRepositoryImpl {
	db.AutoMigrate(&dao.UserDetail{})
	return &UserDetailRepositoryImpl{
		db: db,
	}
}
