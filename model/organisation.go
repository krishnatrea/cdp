package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null"`
	Projects []Project `gorm:"foreignKey:OrgID;constraint:OnDelete:CASCADE"`
	Users    []User    `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
}
