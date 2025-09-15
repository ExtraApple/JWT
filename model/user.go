package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username"`
	Password string `json:"-"`
	Role     string `gorm:"size:32;default:user" json:"role"` // user or admin
}
