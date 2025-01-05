package model

import (
	"gorm.io/gorm"
)

type CustomerData struct {
	gorm.Model
	OrgID     string       `gorm:"type:uuid;not null;index"`
	ProjectID string       `gorm:"type:uuid;not null;index"`
	Data      JSONB        `gorm:"type:jsonb;not null"`
	Org       Organization `gorm:"foreignKey:OrgID;constraint:OnDelete:CASCADE"`
	Project   Project      `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}
