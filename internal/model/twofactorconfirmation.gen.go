// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTwoFactorConfirmation = "TwoFactorConfirmation"

// TwoFactorConfirmation mapped from table <TwoFactorConfirmation>
type TwoFactorConfirmation struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	UserID string `gorm:"column:userId;not null" json:"userId"`
}

// TableName TwoFactorConfirmation's table name
func (*TwoFactorConfirmation) TableName() string {
	return TableNameTwoFactorConfirmation
}
