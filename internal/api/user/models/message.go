package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SenderID   uuid.UUID `gorm:"type:uuid;not null"`
	Sender     User      `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ReceiverID uuid.UUID `gorm:"type:uuid;not null"`
	Receiver   User      `gorm:"foreignKey:ReceiverID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Content    string    `gorm:"not null"`
	SendDate   time.Time `gorm:"not null;default:current_timestamp"`
}
