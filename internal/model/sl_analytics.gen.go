// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSlAnalytic = "sl_analytics"

// SlAnalytic mapped from table <sl_analytics>
type SlAnalytic struct {
	SlName        string `gorm:"column:sl_name;not null" json:"sl_name"`
	Sessions      string `gorm:"column:sessions" json:"sessions"`
	TotalSessions int32  `gorm:"column:total_sessions;not null" json:"total_sessions"`
}

// TableName SlAnalytic's table name
func (*SlAnalytic) TableName() string {
	return TableNameSlAnalytic
}
