package group

import (
	"cmd/gate/main.go/internal/entity"
	"time"
)

type Group struct {
	ID          int            `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt   time.Time      `json:"deleted_at"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	OwnerID     int            `gorm:"index;not null" json:"owner_id"`
	Description string         `json:"description"`
	Users       []*entity.User `gorm:"many2many:user_groups;" json:"users"`
}
