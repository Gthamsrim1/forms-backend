package models

import (
	"time"

	"github.com/google/uuid"
)

type Form struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string     `gorm:"not null"`
	OwnerID   uuid.UUID  `gorm:"type:uuid;not null"`
	Questions []Question `gorm:"foreignKey:FormID"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

type Question struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FormID      uuid.UUID `gorm:"type:uuid;not null"`
	Text        string    `gorm:"not null"`
	Type        string    `gorm:"not null"`
	IsRequired  bool      `gorm:"default:false"`
	OptionsJSON string    `gorm:"type:jsonb"`
	Order       int       `gorm:"default:0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
