package repository

import (
	"context"
	"fmt"
	"log"

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

	result := urp.db.Create(&user)

	if result.Error != nil {

		return fmt.Errorf("create user error: %v", result.Error)
	}
	log.Println("User created successfully:", user)

	return nil
}

func (urp userRepository) Fetch(c context.Context) ([]*model.User, error) {
	var users []*model.User
	err := urp.db.Model(&model.User{}).Find(&users).Error
	if err != nil {

		return nil, err
	}
	return users, nil
}

func (urp userRepository) FetchById(c context.Context, id string) (*model.User, error) {
	user := new(model.User)

	err := urp.db.Model(&model.User{}).First(user, id).Error
	if err != nil {

		return nil, err
	}

	return user, nil
}

func (urp userRepository) FetchByEmail(c context.Context, email string) (model.User, error) {
	return model.User{}, nil
}
