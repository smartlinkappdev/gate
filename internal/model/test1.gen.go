// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTest1 = "test1"

// Test1 mapped from table <test1>
type Test1 struct {
	Date  string `gorm:"column:date;not null" json:"date"`
	Os    string `gorm:"column:os;not null" json:"os"`
	URL   string `gorm:"column:url;not null" json:"url"`
	Users int32  `gorm:"column:users;not null" json:"users"`
	ID    int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
}

// TableName Test1's table name
func (*Test1) TableName() string {
	return TableNameTest1
}
