package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID     string        `gorm:"type:uuid;not null;index"`
	RoleID     string        `gorm:"type:uuid;not null;index"`
	OrgID      *string       `gorm:"type:uuid;index"`
	ProjectID  *string       `gorm:"type:uuid;index"`
	AssignedAt time.Time     `gorm:"autoCreateTime"`
	User       User          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role       Role          `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Org        *Organization `gorm:"foreignKey:OrgID;constraint:OnDelete:SET NULL"`
	Project    *Project      `gorm:"foreignKey:ProjectID;constraint:OnDelete:SET NULL"`
}
