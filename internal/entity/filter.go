package entity

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Filter struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	Description string         `json:"description"`
	//PresetID    int             `json:"owner_id"`
	//Preset      *Preset         `gorm:"foreignKey:PresetID" json:"preset"`
	Data    json.RawMessage `json:"data"`
	Presets []*Preset       `json:"presets" gorm:"many2many:preset_filters;"`
}
