// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameYmData20240425 = "ym_data_2024_04_25"

// YmData20240425 mapped from table <ym_data_2024_04_25>
type YmData20240425 struct {
	Timestamp       time.Time `gorm:"column:timestamp;primaryKey" json:"timestamp"`
	URLPath         string    `gorm:"column:url_path;primaryKey" json:"url_path"`
	Browser         string    `gorm:"column:browser;primaryKey" json:"browser"`
	Device          string    `gorm:"column:device;primaryKey" json:"device"`
	UtmSource       string    `gorm:"column:utm_source;primaryKey" json:"utm_source"`
	UtmMedium       string    `gorm:"column:utm_medium;primaryKey" json:"utm_medium"`
	UtmCampaign     string    `gorm:"column:utm_campaign;primaryKey" json:"utm_campaign"`
	OperatingSystem string    `gorm:"column:operating_system;primaryKey" json:"operating_system"`
	Pageviews       int32     `gorm:"column:pageviews;not null" json:"pageviews"`
	Users           int32     `gorm:"column:users;not null" json:"users"`
}

// TableName YmData20240425's table name
func (*YmData20240425) TableName() string {
	return TableNameYmData20240425
}
