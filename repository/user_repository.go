package repository

import (
	"context"

	"github.com/krishnatrea/cdp/database/model"
	"github.com/krishnatrea/cdp/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return userRepository{db: db}
}

func (urp userRepository) Create(c context.Context, user *model.User) error {
	return nil
}

func (urp userRepository) Fetch(c context.Context) ([]model.User, error) {
	return nil, nil
}

func (urp userRepository) FetchById(c context.Context, id string) (model.User, error) {
	return model.User{}, nil
}

func (urp userRepository) FetchByEmail(c context.Context, email string) (model.User, error) {
	return model.User{}, nil
}
