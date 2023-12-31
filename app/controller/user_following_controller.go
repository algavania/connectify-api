package controller

import (
	service "example/connectify/app/service"

	"github.com/gin-gonic/gin"
)

type UserFollowingController interface {
	GetUserFollowing(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollowingCount(c *gin.Context)
	GetUserFollowersCount(c *gin.Context)
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	CheckHasFollowed(c *gin.Context)
}

type UserFollowingControllerImpl struct {
	svc service.UserFollowingService
}

func (u UserFollowingControllerImpl) GetUserFollowing(c *gin.Context) {
	u.svc.GetUserFollowing(c)
}

func (u UserFollowingControllerImpl) CheckHasFollowed(c *gin.Context) {
	u.svc.CheckHasFollowed(c)
}

func (u UserFollowingControllerImpl) GetUserFollowers(c *gin.Context) {
	u.svc.GetUserFollowers(c)
}

func (u UserFollowingControllerImpl) GetUserFollowingCount(c *gin.Context) {
	u.svc.GetUserFollowingCount(c)
}

func (u UserFollowingControllerImpl) GetUserFollowersCount(c *gin.Context) {
	u.svc.GetUserFollowersCount(c)
}

func (u UserFollowingControllerImpl) Follow(c *gin.Context) {
	u.svc.Follow(c)
}

func (u UserFollowingControllerImpl) Unfollow(c *gin.Context) {
	u.svc.Unfollow(c)
}

func UserFollowingControllerInit(userService service.UserFollowingService) *UserFollowingControllerImpl {
	return &UserFollowingControllerImpl{
		svc: userService,
	}
}
