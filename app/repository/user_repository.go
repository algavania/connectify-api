package repository

import (
	dao "example/connectify/app/domain/dao/user"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserById(id int) (dao.User, error)
	FindUserByEmail(email string) (dao.User, error)
	FindUserByUsername(username string) (dao.User, error)
	Save(user *dao.User) (dao.User, error)
	DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
	db             *gorm.DB
	userDetailRepo UserDetailRepository
}

func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
	user := dao.User{
		ID: id,
	}
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) FindUserByEmail(email string) (dao.User, error) {
	user := dao.User{
		Email: email,
	}
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by email. Error: ", err)
		return dao.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) FindUserByUsername(username string) (dao.User, error) {
	user := dao.User{
		Username: username,
	}
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by username. Error: ", err)
		return dao.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	data, err := u.FindUserByEmail(user.Email)
	log.Info("error in save: ", data.Email)
	if err != nil {
		err = u.db.Create(user).Error
		userDetail := dao.UserDetail{
			Name:   user.Username,
			UserID: user.ID,
		}
		u.userDetailRepo.Save(&userDetail)
	} else {
		user.CreatedAt = data.CreatedAt
		err = u.db.Updates(user).Error
	}
	if err != nil {
		log.Error("Got an error when saving user. Error: ", err)
		return dao.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Unscoped().Delete(&dao.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB, userDetailRepo UserDetailRepository) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db:             db,
		userDetailRepo: userDetailRepo,
	}
}
