package entity

import "time"

type Metric struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Timestamp       time.Time `gorm:"column:timestamp;primaryKey" json:"timestamp"`
	URLPath         string    `gorm:"column:url_path" json:"url_path"`
	Browser         string    `gorm:"column:browser" json:"browser"`
	Device          string    `gorm:"column:device" json:"device"`
	OperatingSystem string    `gorm:"column:operating_system" json:"operating_system"`
	PageViews       int32     `gorm:"column:page_views" json:"page_views"`
	Users           int32     `gorm:"column:users" json:"users"`
}

type YmDatum struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Timestamp       time.Time `gorm:"column:timestamp;primaryKey" json:"timestamp"`
	URLPath         string    `gorm:"column:url_path" json:"url_path"`
	Browser         string    `gorm:"column:browser" json:"browser"`
	Device          string    `gorm:"column:device" json:"device"`
	OperatingSystem string    `gorm:"column:operating_system" json:"operating_system"`
	Pageviews       int32     `gorm:"column:pageviews" json:"pageviews"`
	Users           int32     `gorm:"column:users" json:"users"`
}
