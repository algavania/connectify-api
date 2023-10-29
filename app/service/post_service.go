package service

import (
	"example/connectify/app/constant"
	dao "example/connectify/app/domain/dao/post"
	"example/connectify/app/pkg"
	repository "example/connectify/app/repository"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5/pgconn"
	log "github.com/sirupsen/logrus"
)

type PostService interface {
	GetPostById(c *gin.Context)
	AddPostData(c *gin.Context)
	UpdatePostData(c *gin.Context)
	DeletePost(c *gin.Context)
}

type PostServiceImpl struct {
	postRepository repository.PostRepository
}

func (u PostServiceImpl) UpdatePostData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update post data by id")
	postID, _ := strconv.Atoi(c.Param("postID"))

	var request dao.Post
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.postRepository.FindPostById(postID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	file, err := c.FormFile("media1")
	file2, err2 := c.FormFile("media2")
	file3, err3 := c.FormFile("media3")
	file4, err4 := c.FormFile("media4")

	if err == nil {
		data.Media1 = UploadFile(file, c)
	}

	if err2 == nil {
		data.Media2 = UploadFile(file2, c)
	}

	if err3 == nil {
		data.Media3 = UploadFile(file3, c)
	}

	if err4 == nil {
		data.Media4 = UploadFile(file4, c)
	}

	data.Content = c.PostForm("content")
	res, err := u.postRepository.Save(&data)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if pkg.HandleError(err.(*pgconn.PgError), c) {
			return
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (u PostServiceImpl) GetPostById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get post by id")
	postID, _ := strconv.Atoi(c.Param("postID"))

	data, err := u.postRepository.FindPostById(postID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u PostServiceImpl) AddPostData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data post")
	var request dao.Post
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	file, err := c.FormFile("media1")
	file2, err2 := c.FormFile("media2")
	file3, err3 := c.FormFile("media3")
	file4, err4 := c.FormFile("media4")

	if err == nil {
		request.Media1 = UploadFile(file, c)
	}

	if err2 == nil {
		request.Media2 = UploadFile(file2, c)
	}

	if err3 == nil {
		request.Media3 = UploadFile(file3, c)
	}

	if err4 == nil {
		request.Media4 = UploadFile(file4, c)
	}

	request.Content = c.PostForm("content")
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	request.UserID = userID
	parentPostID, _ := strconv.Atoi(c.PostForm("parent_post_id"))
	if parentPostID == 0 {
		request.ParentPostID = nil
	} else {
		request.ParentPostID = &parentPostID
	}

	data, err := u.postRepository.Save(&request)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if pkg.HandleError(err.(*pgconn.PgError), c) {
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func UploadFile(file *multipart.FileHeader, c *gin.Context) string {
	if !pkg.IsImageFile(file) && !pkg.IsVideoFile(file) {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse_(constant.InvalidRequest.GetResponseStatus(), "File is not image or video", pkg.Null()))
	}
	// Save the uploaded file to the server
	url := "public/media/" + strconv.FormatInt(time.Now().UTC().UnixMilli(), 10) + filepath.Ext(file.Filename)
	err := c.SaveUploadedFile(file, url)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}
	return url
}

func (u PostServiceImpl) DeletePost(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data post by id")
	postID, _ := strconv.Atoi(c.Param("postID"))

	err := u.postRepository.DeletePostById(postID)
	if err != nil {
		log.Error("Error happened when try delete data post from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func PostServiceInit(postRepository repository.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{
		postRepository: postRepository,
	}
}
