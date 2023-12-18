package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model

	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(255)"`
	Type      string    `gorm:"type:varchar(255)"`
	Region    string    `gorm:"type:varchar(255)"`
	Image     string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
