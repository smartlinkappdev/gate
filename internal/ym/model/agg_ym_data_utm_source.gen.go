// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAggYmDataUtmSource = "agg_ym_data_utm_source"

// AggYmDataUtmSource mapped from table <agg_ym_data_utm_source>
type AggYmDataUtmSource struct {
	Timestamp time.Time `gorm:"column:timestamp;primaryKey" json:"timestamp"`
	UtmSource string    `gorm:"column:utm_source;primaryKey" json:"utm_source"`
	Pageviews int32     `gorm:"column:pageviews" json:"pageviews"`
	Users     int32     `gorm:"column:users" json:"users"`
}

// TableName AggYmDataUtmSource's table name
func (*AggYmDataUtmSource) TableName() string {
	return TableNameAggYmDataUtmSource
}
