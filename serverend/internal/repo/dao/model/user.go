package model

import (
	"database/sql"
)

type User struct {
	Id        int64          `gorm:"column:id;primary_key;auto_increment"`
	Username  string         `gorm:"column:username;unique"`
	Password  string         `gorm:"column:password"`
	Phone     sql.NullString `gorm:"column:phone;unique"`
	Email     sql.NullString `gorm:"column:email;unique"`
	CreatedAt int64          `gorm:"column:created_at"`
	UpdatedAt int64          `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}
