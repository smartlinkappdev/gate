package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Role      string         `json:"role"`
	Groups    []*Group       `json:"groups" gorm:"many2many:user_groups;"`
}
