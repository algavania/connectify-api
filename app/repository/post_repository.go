package repository

import (
	dao "example/connectify/app/domain/dao/post"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindPostById(id int) (dao.Post, error)
	Save(post *dao.Post) (dao.Post, error)
	DeletePostById(id int) error
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

func (u PostRepositoryImpl) FindPostById(id int) (dao.Post, error) {
	post := dao.Post{
		ID: id,
	}
	err := u.db.First(&post).Error
	if err != nil {
		log.Error("Got and error when find post by id. Error: ", err)
		return dao.Post{}, err
	}
	return post, nil
}
func (u PostRepositoryImpl) Save(post *dao.Post) (dao.Post, error) {

	data, err := u.FindPostById(post.ID)
	if err != nil {
		err = u.db.Create(post).Error
	} else {
		post.CreatedAt = data.CreatedAt
		err = u.db.Updates(post).Error
	}
	if err != nil {
		log.Error("Got an error when saving post. Error: ", err)
		return dao.Post{}, err
	}
	return *post, nil
}

func (u PostRepositoryImpl) DeletePostById(id int) error {
	err := u.db.Unscoped().Delete(&dao.Post{}, id).Error
	if err != nil {
		log.Error("Got an error when delete post. Error: ", err)
		return err
	}
	return nil
}

func PostRepositoryInit(db *gorm.DB) *PostRepositoryImpl {
	db.AutoMigrate(&dao.Post{})
	return &PostRepositoryImpl{
		db: db,
	}
}
