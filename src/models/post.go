package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title     string
	Body      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	AuthorID  string    `gorm:"references:ID"`
	Author    User
}
