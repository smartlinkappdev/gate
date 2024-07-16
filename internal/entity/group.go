package entity

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	OwnerID     int            `json:"owner_id"`
	Owner       *User          `gorm:"foreignKey:OwnerID"`
	Description string         `json:"description"`
	Users       []*User        `gorm:"many2many:user_groups;" json:"users"`
	Links       []*Link        `gorm:"many2many:group_link;" json:"links"`
}
