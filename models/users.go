package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email            string    `gorm:"type:citext;uniqueIndex;not null"`
	PasswordHash     string    `gorm:"not null"`
	IsActive         bool      `gorm:"default:true"`
	FailedLoginCount int       `gorm:"default:0"`
	LastFailedLogin  *time.Time
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	RefreshTokens    []RefreshToken `gorm:"foreignKey:UserID"`
}

type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	TokenHash string    `gorm:"not null;index"`
	ExpiresAt time.Time `gorm:"not null"`
	RevokedAt *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
