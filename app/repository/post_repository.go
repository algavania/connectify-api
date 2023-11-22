package repository

import (
	dao "example/connectify/app/domain/dao/post"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindAllPosts() ([]dao.PostResponse, error)
	FindAllReplies(id int) ([]dao.Post, error)
	FindPostById(id int) (dao.Post, error)
	Save(post *dao.Post) (dao.Post, error)
	DeletePostById(id int) error
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

func (u PostRepositoryImpl) FindAllPosts() ([]dao.PostResponse, error) {
	var posts []dao.Post
	var postsRes []dao.PostResponse
	err := u.db.Preload("User").Where("parent_post_id IS NULL").Order("created_at desc").Find(&posts).Error
	if err != nil {
		log.Error("Got and error when find post by id. Error: ", err)
		return []dao.PostResponse{}, err
	}

	for i, post := range posts {
		var commentCount int64
		err := u.db.Model(&dao.Post{}).Where("parent_post_id = ?", post.ID).Count(&commentCount).Error
		if err != nil {
			log.Error("Got an error when counting comments. Error: ", err)
			return []dao.PostResponse{}, err
		}
		postsRes = append(postsRes, dao.PostResponse{Post: post})
		postsRes[i].CommentCount = commentCount
	}
	return postsRes, nil
}

func (u PostRepositoryImpl) FindAllReplies(id int) ([]dao.Post, error) {
	var posts []dao.Post
	err := u.db.Preload("User").Preload("ParentPost").Where("parent_post_id = ?", id).Order("created_at desc").Find(&posts).Error
	if err != nil {
		log.Error("Got and error when find replies. Error: ", err)
		return []dao.Post{}, err
	}
	return posts, nil
}

func (u PostRepositoryImpl) FindPostById(id int) (dao.Post, error) {
	post := dao.Post{
		ID: id,
	}
	err := u.db.Preload("User").Where("id = ?", id).First(&post).Error
	if err != nil {
		log.Error("Got and error when find post by id. Error: ", err)
		return dao.Post{}, err
	}
	return post, nil
}

func (u PostRepositoryImpl) Save(post *dao.Post) (dao.Post, error) {
	log.Info("post id: ", post.ID)
	data, err := u.FindPostById(post.ID)
	log.Info("data id: ", data.ID)
	if err != nil {
		err = u.db.Create(post).Error
		postRes, _ := u.FindPostById(post.ID)
		post = &postRes
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
