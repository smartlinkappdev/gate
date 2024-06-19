package entity

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OwnerID   int            `json:"owner_id"`
	Owner     *User          `gorm:"foreignKey:OwnerID"`
}
