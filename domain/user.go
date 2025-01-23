package domain

import (
	"context"

	"github.com/krishnatrea/cdp/database/model"
)

type UserRepository interface {
	Create(c context.Context, user *model.User) error
	Fetch(c context.Context) ([]*model.User, error)
	FetchById(c context.Context, id string) (*model.User, error)
	FetchByEmail(c context.Context, email string) (model.User, error)
}
