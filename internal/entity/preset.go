package entity

import (
	"gorm.io/gorm"
	"time"
)

type Preset struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	Description string         `json:"description"`

	OwnerID *int  `json:"owner_id"`
	Owner   *User `gorm:"foreignKey:OwnerID" json:"owner"`

	GroupID *int   `json:"group_id"`
	Group   *Group `gorm:"foreignKey:GroupID" json:"group"`

	LinkID *int  `json:"link_id"`
	Link   *Link `gorm:"foreignKey:LinkID" json:"link"`

	FilterID *int    `json:"filter_id"`
	Filter   *Filter `gorm:"foreignKey:FilterID" json:"filter"`

	Users   []*User   `gorm:"many2many:preset_users;" json:"users"`
	Links   []*Link   `gorm:"many2many:preset_links;" json:"links"`
	Groups  []*Group  `gorm:"many2many:preset_groups;" json:"groups"`
	Filters []*Filter `gorm:"many2many:preset_filters;" json:"filters"`
}
