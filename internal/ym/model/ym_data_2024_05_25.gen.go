// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameYmData20240525 = "ym_data_2024_05_25"

// YmData20240525 mapped from table <ym_data_2024_05_25>
type YmData20240525 struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Timestamp       time.Time `gorm:"column:timestamp;primaryKey" json:"timestamp"`
	URLPath         string    `gorm:"column:url_path" json:"url_path"`
	Browser         string    `gorm:"column:browser" json:"browser"`
	Device          string    `gorm:"column:device" json:"device"`
	OperatingSystem string    `gorm:"column:operating_system" json:"operating_system"`
	Pageviews       int32     `gorm:"column:pageviews" json:"pageviews"`
	Users           int32     `gorm:"column:users" json:"users"`
}

// TableName YmData20240525's table name
func (*YmData20240525) TableName() string {
	return TableNameYmData20240525
}
