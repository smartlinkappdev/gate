package entity

import (
	"gorm.io/gorm"
	"time"
)

type Feedback struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OwnerID   int            `json:"owner_id"`
	Owner     *User          `gorm:"foreignKey:OwnerID"`
	Active    bool           `gorm:"default:true" json:"active"`
	Type      string         `json:"type"`
	Title     string         `gorm:"size:255" json:"title"`
	Content   string         `json:"content"`
	Status    string         `json:"status"`
}
