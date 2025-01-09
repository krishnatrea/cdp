package model

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleID       string     `gorm:"type:uuid;not null;index"`
	PermissionID string     `gorm:"type:uuid;not null;index"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
}
