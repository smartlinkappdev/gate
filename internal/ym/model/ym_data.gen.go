// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameYmDatum = "ym_data"

// YmDatum mapped from table <ym_data>
type YmDatum struct {
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

// TableName YmDatum's table name
func (*YmDatum) TableName() string {
	return TableNameYmDatum
}
