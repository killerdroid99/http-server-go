package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	Name     string
	Email    string `gorm:"index:idx_name,unique"`
	Password string
	Created  time.Time `gorm:"autoCreateTime"`
}
