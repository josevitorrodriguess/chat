package service

import (
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"github.com/josevitorrodriguess/chat/internal/api/user/repository"
)


type MessageService interface {
	Create(*models.Message) (models.Message, error)
	GetByID(id uint) (models.Message, error)
	GetAll() ([]models.Message, error)
	Update(*models.Message) (models.Message, error)
	Delete(id uint) error
}

type messageService struct {
	repo repository.MessageRepository
}


func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{
		repo: repo,
	}
}


func (ms *messageService) Create(message *models.Message) (models.Message, error) {
	return ms.repo.Create(message)
}


func (ms *messageService) GetByID(id uint) (models.Message, error) {
	return ms.repo.GetByID(id)
}

func (ms *messageService) GetAll() ([]models.Message, error) {
	return ms.repo.GetAll()
}


func (ms *messageService) Update(message *models.Message) (models.Message, error) {
	return ms.repo.Update(message)
}


func (ms *messageService) Delete(id uint) error {
	return ms.repo.Delete(id)
}
