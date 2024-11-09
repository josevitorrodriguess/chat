package message

import (
	"time"

	"github.com/google/uuid"
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
)

type Message struct {
	ID         uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SenderID   uuid.UUID   `gorm:"type:uuid;not null"` 
	Sender     models.User `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ReceiverID uuid.UUID   `gorm:"type:uuid;not null"` 
	Receiver   models.User `gorm:"foreignKey:ReceiverID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Content    string      `gorm:"not null"`
	SendDate   time.Time   `gorm:"not null;default:current_timestamp"`
}
