package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name           string       `gorm:"type:varchar(255);not null"`
	IsCustomizable bool         `gorm:"default:true"`
	Permissions    []Permission `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE"`
}
