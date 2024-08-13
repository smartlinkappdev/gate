package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Role      string         `json:"role"`
	Groups    []*Group       `json:"groups" gorm:"many2many:user_groups;"`
	Links     []*Link        `json:"links" gorm:"many2many:user_link;"`
	Presets   []*Preset      `json:"presets" gorm:"many2many:preset_users;"`
}
