package domain

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID                uuid.UUID
	Title             string
	Content           string
	CreatedAt         time.Time
	CommentPermission bool
}

type Comment struct {
	ID        uuid.UUID
	Content   string
	CreatedAt time.Time
	ArticleID uuid.UUID
	ParentID  uuid.UUID
}
