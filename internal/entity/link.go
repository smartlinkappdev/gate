package entity

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID          int `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name" gorm:"name"`
	Path        string         `json:"path" gorm:"path"`
	Description string         `json:"description" gorm:"description"`
	Image       string         `json:"image" gorm:"image"`
	OwnerID     int            `json:"owner_id"`
	Owner       *User          `gorm:"foreignKey:OwnerID"`
	Groups      []*Group       `gorm:"many2many:group_link;" json:"groups"`
	Users       []*User        `gorm:"many2many:user_link;" json:"users"`
	Presets     []*Preset      `json:"presets" gorm:"many2many:preset_links"`
}
