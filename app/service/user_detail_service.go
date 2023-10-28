package service

import (
	"example/connectify/app/constant"
	dao "example/connectify/app/domain/dao/user"
	"example/connectify/app/pkg"
	repository "example/connectify/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	log "github.com/sirupsen/logrus"
)

type UserDetailService interface {
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
}

type UserDetailServiceImpl struct {
	UserDetailRepository repository.UserDetailRepository
}

func (u UserDetailServiceImpl) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.UserDetailRepository.FindUserById(userID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserDetailServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request dao.UserDetail
	userID, _ := strconv.Atoi(c.Param("userID"))
	request.UserID = userID

	err := c.Request.ParseForm()
	if err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	request.Name = c.PostForm("name")
	request.Description = c.PostForm("description")
	request.Birthday, err = pkg.ParseTime(c.PostForm("birthday"))
	if err != nil {
		log.Error("Error happened in date", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	file, err := c.FormFile("file")

	if err == nil {
		// Save the uploaded file to the server
		url := "public/images/user/" + file.Filename
		err = c.SaveUploadedFile(file, url)
		if err != nil {
			pkg.PanicException(constant.UnknownError)
		}
		log.Info("file name " + file.Filename)

		request.PhotoUrl = url
	}

	data, err := u.UserDetailRepository.Save(&request)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if err != nil {
			if pkg.HandleError(err.(*pgconn.PgError), c) {
				return
			}
		}
	}
	log.Info("request ", request.Name, request.Description, request.Birthday, request.PhotoUrl)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func UserDetailServiceInit(UserDetailRepository repository.UserDetailRepository) *UserDetailServiceImpl {
	return &UserDetailServiceImpl{
		UserDetailRepository: UserDetailRepository,
	}
}
