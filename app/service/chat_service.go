package service

import (
	"example/connectify/app/constant"
	dao "example/connectify/app/domain/dao/chat"
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

type ChatService interface {
	GetChatById(c *gin.Context)
	AddChatData(c *gin.Context)
	UpdateChatData(c *gin.Context)
	AddMessage(c *gin.Context)
	AddParticipant(c *gin.Context)
	DeleteMessage(c *gin.Context)
	DeleteParticipant(c *gin.Context)
	DeleteChat(c *gin.Context)
}

type ChatServiceImpl struct {
	chatRepository repository.ChatRepository
}

func (u ChatServiceImpl) GetChatById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get chat by id")
	chatID, _ := strconv.Atoi(c.Param("chatID"))

	data, err := u.chatRepository.FindChatById(chatID)
	if err != nil {
		log.Error("Error happened when getting data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ChatServiceImpl) AddChatData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data chat")
	var request dao.Chat
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.chatRepository.SaveChat(&request)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if pkg.HandleError(err.(*pgconn.PgError), c) {
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ChatServiceImpl) UpdateChatData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data chat")
	ChatID, _ := strconv.Atoi(c.PostForm("chat_id"))

	var request dao.Chat
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	request.ID = ChatID

	data, err := u.chatRepository.SaveChat(&request)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if pkg.HandleError(err.(*pgconn.PgError), c) {
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ChatServiceImpl) AddParticipant(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data chat")
	var request dao.ChatParticipant
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.chatRepository.AddParticipant(&request)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if pkg.HandleError(err.(*pgconn.PgError), c) {
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ChatServiceImpl) AddMessage(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data chat")
	var request dao.ChatMessage
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		log.Error("Error happened when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	file, err := c.FormFile("file")

	if err == nil {
		request.Media = UploadMediaFile(file, c)
	}

	request.Content = c.PostForm("content")

	data, err := u.chatRepository.SaveMessage(&request)
	if err != nil {
		log.Error("Error happened when saving data to database. Error", err)
		if pkg.HandleError(err.(*pgconn.PgError), c) {
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func UploadMediaFile(file *multipart.FileHeader, c *gin.Context) string {
	// Save the uploaded file to the server
	url := "public/media/" + strconv.FormatInt(time.Now().UTC().UnixMilli(), 10) + filepath.Ext(file.Filename)
	err := c.SaveUploadedFile(file, url)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}
	return url
}

func (u ChatServiceImpl) DeleteChat(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data chat by id")
	chatID, _ := strconv.Atoi(c.Param("chatID"))

	err := u.chatRepository.DeleteChat(chatID)
	if err != nil {
		log.Error("Error happened when try delete data chat from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (u ChatServiceImpl) DeleteMessage(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data chat by id")
	chatID, _ := strconv.Atoi(c.Param("messageID"))

	err := u.chatRepository.DeleteMessage(chatID)
	if err != nil {
		log.Error("Error happened when try delete data chat from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}
func (u ChatServiceImpl) DeleteParticipant(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data chat by id")
	chatID, _ := strconv.Atoi(c.Param("participantID"))

	err := u.chatRepository.DeleteParticipant(chatID)
	if err != nil {
		log.Error("Error happened when try delete data chat from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}
func ChatServiceInit(chatRepository repository.ChatRepository) *ChatServiceImpl {
	return &ChatServiceImpl{
		chatRepository: chatRepository,
	}
}
