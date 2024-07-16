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
	Name      string         `json:"name" gorm:"name"`
	OwnerID   int            `json:"owner_id"`
	Owner     *User          `gorm:"foreignKey:OwnerID"`
	Groups    []*Group       `gorm:"many2many:group_link;" json:"groups"`
	Users     []*User        `gorm:"many2many:user_link;" json:"users"`
}
