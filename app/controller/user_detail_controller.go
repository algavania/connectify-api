package controller

import (
	service "example/connectify/app/service"

	"github.com/gin-gonic/gin"
)

type UserDetailController interface {
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserDetailControllerImpl struct {
	svc service.UserDetailService
}

func (u UserDetailControllerImpl) GetAllUsers(c *gin.Context) {
	u.svc.GetAllUsers(c)
}

func (u UserDetailControllerImpl) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserDetailControllerImpl) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

func UserDetailControllerInit(userService service.UserDetailService) *UserDetailControllerImpl {
	return &UserDetailControllerImpl{
		svc: userService,
	}
}
