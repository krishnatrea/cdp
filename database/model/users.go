package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Firstname string `gorm:"type:varchar(255);"`
	Lastname  string `gorm:"type:varchar(255);"`
	Password  string `grom:"type:varchar(255);"`
	IsDeleted bool   `grom:"type:boolean;"`
	Roles     []Role `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
}
