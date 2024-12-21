package repository

import (
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"gorm.io/gorm"
)


type MessageRepository interface {
	Create(*models.Message) (models.Message, error)
	GetByID(id uint) (models.Message, error)
	GetAll() ([]models.Message, error)
	Update(*models.Message) (models.Message, error)
	Delete(id uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}


func (mr *messageRepository) Create(message *models.Message) (models.Message, error) {
	if err := mr.db.Create(message).Error; err != nil {
		return models.Message{}, err
	}
	return *message, nil
}

func (mr *messageRepository) GetByID(id uint) (models.Message, error) {
	var message models.Message
	if err := mr.db.First(&message, id).Error; err != nil {
		return models.Message{}, err
	}
	return message, nil
}

func (mr *messageRepository) GetAll() ([]models.Message, error) {
	var messages []models.Message
	if err := mr.db.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}


func (mr *messageRepository) Update(message *models.Message) (models.Message, error) {
	if err := mr.db.Save(message).Error; err != nil {
		return models.Message{}, err
	}
	return *message, nil
}


func (mr *messageRepository) Delete(id uint) error {
	if err := mr.db.Delete(&models.Message{}, id).Error; err != nil {
		return err
	}
	return nil
}
