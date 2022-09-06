package models

import (
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:60;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;unique" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp(3)" json:"updated_at"`
}
