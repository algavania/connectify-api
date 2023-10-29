package repository

import (
	dao "example/connectify/app/domain/dao/chat"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ChatRepository interface {
	FindChatById(id int) (dao.Chat, error)
	SaveMessage(chat *dao.ChatMessage) (dao.ChatMessage, error)
	SaveChat(chat *dao.Chat) (dao.Chat, error)
	AddParticipant(chat *dao.ChatParticipant) (dao.ChatParticipant, error)
	DeleteParticipant(id int) error
	DeleteMessage(id int) error
	DeleteChat(id int) error
}

type ChatRepositoryImpl struct {
	db *gorm.DB
}

func (u ChatRepositoryImpl) FindChatById(id int) (dao.Chat, error) {
	var chat dao.Chat
	if err := u.db.Preload("Messages").Preload("Participants").First(&chat, id).Error; err != nil {
		log.Error("Got an error when finding chat by id. Error: ", err)
		return dao.Chat{}, err
	}
	return chat, nil
}
func (u ChatRepositoryImpl) SaveMessage(chat *dao.ChatMessage) (dao.ChatMessage, error) {

	data, err := u.FindChatById(chat.ID)
	if err != nil {
		err = u.db.Create(chat).Error
	} else {
		chat.CreatedAt = data.CreatedAt
		err = u.db.Updates(chat).Error
	}
	if err != nil {
		log.Error("Got an error when saving chat. Error: ", err)
		return dao.ChatMessage{}, err
	}
	return *chat, nil
}

func (u ChatRepositoryImpl) SaveChat(chat *dao.Chat) (dao.Chat, error) {

	data, err := u.FindChatById(chat.ID)
	if err != nil {
		err = u.db.Create(chat).Error
	} else {
		chat.CreatedAt = data.CreatedAt
		err = u.db.Updates(chat).Error
	}
	if err != nil {
		log.Error("Got an error when saving chat. Error: ", err)
		return dao.Chat{}, err
	}
	return *chat, nil
}

func (u ChatRepositoryImpl) AddParticipant(chat *dao.ChatParticipant) (dao.ChatParticipant, error) {

	data, err := u.FindChatById(chat.ID)
	if err != nil {
		err = u.db.Create(chat).Error
	} else {
		chat.CreatedAt = data.CreatedAt
		err = u.db.Updates(chat).Error
	}
	if err != nil {
		log.Error("Got an error when saving chat. Error: ", err)
		return dao.ChatParticipant{}, err
	}
	return *chat, nil
}

func (u ChatRepositoryImpl) DeleteParticipant(id int) error {
	err := u.db.Unscoped().Delete(&dao.ChatParticipant{}, id).Error
	if err != nil {
		log.Error("Got an error when delete perticipant. Error: ", err)
		return err
	}
	return nil
}

func (u ChatRepositoryImpl) DeleteMessage(id int) error {
	err := u.db.Unscoped().Delete(&dao.ChatMessage{}, id).Error
	if err != nil {
		log.Error("Got an error when delete chat message. Error: ", err)
		return err
	}
	return nil
}

func (u ChatRepositoryImpl) DeleteChat(id int) error {
	err := u.db.Unscoped().Delete(&dao.Chat{}, id).Error
	if err != nil {
		log.Error("Got an error when delete chat message. Error: ", err)
		return err
	}
	return nil
}

func ChatRepositoryInit(db *gorm.DB) *ChatRepositoryImpl {
	db.AutoMigrate(&dao.Chat{}, &dao.ChatMessage{}, &dao.ChatParticipant{})
	return &ChatRepositoryImpl{
		db: db,
	}
}
