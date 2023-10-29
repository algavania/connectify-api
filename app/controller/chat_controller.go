package controller

import (
	service "example/connectify/app/service"

	"github.com/gin-gonic/gin"
)

type ChatController interface {
	GetChatById(c *gin.Context)
	AddChatData(c *gin.Context)
	UpdateChatData(c *gin.Context)
	AddMessage(c *gin.Context)
	AddParticipant(c *gin.Context)
	DeleteMessage(c *gin.Context)
	DeleteParticipant(c *gin.Context)
	DeleteChat(c *gin.Context)
}

type ChatControllerImpl struct {
	svc service.ChatService
}

func (u ChatControllerImpl) GetChatById(c *gin.Context) {
	u.svc.GetChatById(c)
}
func (u ChatControllerImpl) AddChatData(c *gin.Context) {
	u.svc.AddChatData(c)
}
func (u ChatControllerImpl) UpdateChatData(c *gin.Context) {
	u.svc.UpdateChatData(c)
}
func (u ChatControllerImpl) AddMessage(c *gin.Context) {
	u.svc.AddMessage(c)
}
func (u ChatControllerImpl) AddParticipant(c *gin.Context) {
	u.svc.AddParticipant(c)
}
func (u ChatControllerImpl) DeleteMessage(c *gin.Context) {
	u.svc.DeleteMessage(c)
}
func (u ChatControllerImpl) DeleteParticipant(c *gin.Context) {
	u.svc.DeleteParticipant(c)
}
func (u ChatControllerImpl) DeleteChat(c *gin.Context) {
	u.svc.DeleteChat(c)
}

func ChatControllerInit(chatService service.ChatService) *ChatControllerImpl {
	return &ChatControllerImpl{
		svc: chatService,
	}
}
