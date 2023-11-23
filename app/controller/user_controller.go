package controller

import (
	service "example/connectify/app/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddUserData(c *gin.Context)
	GetUserByUsername(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
	Login(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func (u UserControllerImpl) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserControllerImpl) GetUserByUsername(c *gin.Context) {
	u.svc.GetUserByUsername(c)
}

func (u UserControllerImpl) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUserData(c)
}

func (u UserControllerImpl) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}

func (u UserControllerImpl) Login(c *gin.Context) {
	u.svc.Login(c)
}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
