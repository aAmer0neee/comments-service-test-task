package repository

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Article struct {
	ID                uuid.UUID `gorm:"primaryKey;type:uuid"`
	Content           string    `gorm:"type:text;not null"`
	CreatedAt         time.Time `gorm:"not null"`
	CommentPermission bool      `gorm:"default:true"`
}

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	Content   string    `gorm:"type:varchar(2000);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	ArticleID uuid.UUID `gorm:"type:uuid;not null;index" `
	ParentID  uuid.UUID `gorm:"type:uuid;index;default:NULL" `
}
