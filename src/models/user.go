package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name      string
	Email     string `gorm:"index:idx_name,unique"`
	Password  string
	ImgUrl    *string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Posts     []Post    `gorm:"foreignKey:AuthorID"`
}
