package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	OrgID        string         `gorm:"type:uuid;not null;index"`
	Name         string         `gorm:"type:varchar(255);not null"`
	Users        []User         `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE"`
	Org          Organization   `gorm:"foreignKey:OrgID;constraint:OnDelete:CASCADE"`
	CustomerData []CustomerData `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}
