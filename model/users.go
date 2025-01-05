package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique;not null"`
	FullName string `gorm:"type:varchar(255);"`
	Roles    []Role `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
}
