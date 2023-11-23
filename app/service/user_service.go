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
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserByUsername(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
	Login(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update user data by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Username = request.Username
	data.Email = request.Email
	data.Password = request.Password
	u.userRepository.Save(&data)

	if err != nil {
		log.Error("Error happened when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetUserByUsername(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	username := c.Param("username")

	data, err := u.userRepository.FindUserByUsername(username)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	email, _ := u.userRepository.FindUserByEmail(request.Email)
	if email.Email != "" {
		pkg.CustomPanicException(http.StatusBadRequest, "Email already exist", c)
		return
	}

	username, _ := u.userRepository.FindUserByUsername(request.Username)
	if username.Username != "" {
		pkg.CustomPanicException(http.StatusBadRequest, "Username already exist", c)
		return
	}

	data, err := u.userRepository.Save(&request)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pkg.HandleError(pgErr, c) {
				return
			}
		} else {
			pkg.PanicException(constant.UnknownError)
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	err := u.userRepository.DeleteUserById(userID)
	if err != nil {
		log.Error("Error happened when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (u UserServiceImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program login user")
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	log.Info("request", request)

	if request.Email == "" || request.Password == "" {
		var title = "Email"
		if request.Password == "" {
			title = "Password"
		}
		pkg.CustomPanicException(http.StatusBadRequest, title+" cannot be empty", c)
		return
	}

	user, emailErr := u.userRepository.FindUserByEmail(request.Email)
	if user.Email == "" {
		pkg.CustomPanicException(http.StatusBadRequest, "Email not found", c)
		return
	}
	if emailErr != nil {
		log.Error("Error happened when getting user from DB. Error:", emailErr)
		pkg.PanicException(constant.UnknownError)
	}

	log.Info("hashed password", user)
	log.Info("request password", request.Password)

	if passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); passwordErr != nil {
		log.Error("Error happened when comparing password. Error:", passwordErr)
		pkg.CustomPanicException(http.StatusBadRequest, "Wrong password", c)
		return
	}

	tokenString, tokenErr := pkg.GenerateToken(user.ID)
	if tokenErr != nil {
		log.Error("Error happened when signing token. Error:", tokenErr)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{
		"token": tokenString,
		"user":  user,
	}))
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
