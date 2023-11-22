package controller

import (
	service "example/connectify/app/service"

	"github.com/gin-gonic/gin"
)

type PostController interface {
	AddPostData(c *gin.Context)
	GetPostById(c *gin.Context)
	GetAllPosts(c *gin.Context)
	GetAllReplies(c *gin.Context)
	UpdatePostData(c *gin.Context)
	DeletePost(c *gin.Context)
}

type PostControllerImpl struct {
	svc service.PostService
}

func (u PostControllerImpl) AddPostData(c *gin.Context) {
	u.svc.AddPostData(c)
}

func (u PostControllerImpl) GetPostById(c *gin.Context) {
	u.svc.GetPostById(c)
}

func (u PostControllerImpl) GetAllPosts(c *gin.Context) {
	u.svc.GetAllPosts(c)
}

func (u PostControllerImpl) GetAllReplies(c *gin.Context) {
	u.svc.GetAllReplies(c)
}

func (u PostControllerImpl) UpdatePostData(c *gin.Context) {
	u.svc.UpdatePostData(c)
}

func (u PostControllerImpl) DeletePost(c *gin.Context) {
	u.svc.DeletePost(c)
}

func PostControllerInit(postService service.PostService) *PostControllerImpl {
	return &PostControllerImpl{
		svc: postService,
	}
}
