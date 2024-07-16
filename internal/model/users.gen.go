// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID                          string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt                   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Password                    string    `gorm:"column:password" json:"password"`
	Salt                        string    `gorm:"column:salt" json:"salt"`
	Username                    string    `gorm:"column:username" json:"username"`
	Email                       string    `gorm:"column:email" json:"email"`
	EmailVerified               time.Time `gorm:"column:emailVerified" json:"emailVerified"`
	IsTwoFactorEnabled          bool      `gorm:"column:isTwoFactorEnabled;not null" json:"isTwoFactorEnabled"`
	Role                        string    `gorm:"column:role;not null;default:USER" json:"role"`
	ResetPasswordEmailTimestamp time.Time `gorm:"column:reset_password_email_timestamp" json:"reset_password_email_timestamp"`
	VerificationEmailTimestamp  time.Time `gorm:"column:verification_email_timestamp" json:"verification_email_timestamp"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
