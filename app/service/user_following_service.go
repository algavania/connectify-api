package service

import (
	"example/connectify/app/constant"
	dao "example/connectify/app/domain/dao/user"
	"example/connectify/app/pkg"
	repository "example/connectify/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserFollowingService interface {
	GetUserFollowing(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollowingCount(c *gin.Context)
	GetUserFollowersCount(c *gin.Context)
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	CheckHasFollowed(c *gin.Context)
}

type UserFollowingServiceImpl struct {
	userFollowingRepository repository.UserFollowingRepository
}

func (u UserFollowingServiceImpl) Follow(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program follow userFollowing data by id")

	var request dao.UserFollowing
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	userId, _ := pkg.ExtractTokenID(c)
	request.FollowedUserID = request.UserID
	request.UserID = int(userId)
	log.Info("request", request)
	data, err := u.userFollowingRepository.Follow(&request)

	if err != nil {
		log.Error("Error happened when follow data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserFollowingServiceImpl) GetUserFollowing(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get userFollowing by id")
	userFollowingID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userFollowingRepository.GetUserFollowing(userFollowingID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserFollowingServiceImpl) GetUserFollowers(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get userFollowing by id")
	userFollowingID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userFollowingRepository.GetUserFollowers(userFollowingID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserFollowingServiceImpl) GetUserFollowingCount(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get userFollowing by id")
	userFollowingID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userFollowingRepository.GetUserFollowingCount(userFollowingID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserFollowingServiceImpl) GetUserFollowersCount(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get userFollowing by id")
	userFollowingID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userFollowingRepository.GetUserFollowersCount(userFollowingID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserFollowingServiceImpl) CheckHasFollowed(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get userFollowing by id")
	userFollowedId, _ := strconv.Atoi(c.Param("userID"))

	userId, _ := pkg.ExtractTokenID(c)

	data := u.userFollowingRepository.CheckHasFollowed(userFollowedId, int(userId))

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserFollowingServiceImpl) Unfollow(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data userFollowing by id")

	userFollowedId, _ := strconv.Atoi(c.Param("userID"))
	userId, _ := pkg.ExtractTokenID(c)

	err := u.userFollowingRepository.Unfollow(userFollowedId, int(userId))
	if err != nil {
		log.Error("Error happened when try delete data userFollowing from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func UserFollowingServiceInit(userFollowingRepository repository.UserFollowingRepository) *UserFollowingServiceImpl {
	return &UserFollowingServiceImpl{
		userFollowingRepository: userFollowingRepository,
	}
}
