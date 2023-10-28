package repository

import (
	dao "example/connectify/app/domain/dao/user"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserDetailRepository interface {
	FindUserById(id int) (dao.UserDetail, error)
	Save(user *dao.UserDetail) (dao.UserDetail, error)
}

type UserDetailRepositoryImpl struct {
	db *gorm.DB
}

func (u UserDetailRepositoryImpl) FindUserById(id int) (dao.UserDetail, error) {
	user := dao.UserDetail{
		UserID: id,
	}
	err := u.db.First(&user).Error
	if err != nil {
		log.Error("Got and error when find user detail by id. Error: ", err)
		return dao.UserDetail{}, err
	}
	return user, nil
}

func (u UserDetailRepositoryImpl) Save(user *dao.UserDetail) (dao.UserDetail, error) {
	var err = u.db.Save(user).Error
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
