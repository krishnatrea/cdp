package repository

import (
	"github.com/krishnatrea/cdp/domain"
	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

// AssignRole implements domain.ProjectRepository.
func (p projectRepository) AssignRole() {
	panic("unimplemented")
}

// Create implements domain.ProjectRepository.
func (p projectRepository) Create() {
	panic("unimplemented")
}

// GetAllByUserId implements domain.ProjectRepository.
func (p projectRepository) GetAllByUserId() {
	panic("unimplemented")
}

// GetByID implements domain.ProjectRepository.
func (p projectRepository) GetByID() {
	panic("unimplemented")
}

func NewProjectRepository(db *gorm.DB) domain.ProjectRepository {
	return projectRepository{
		db: db,
	}
}
