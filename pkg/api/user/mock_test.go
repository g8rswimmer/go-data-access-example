package user

import (
	"context"

	"github.com/g8rswimmer/go-data-access-example/pkg/model"
)

type mockUserDAO struct {
	user  *model.UserEntity
	users []*model.UserEntity
	err   error
}

func (m *mockUserDAO) Create(ctx context.Context, user *model.User) (*model.UserEntity, error) {
	return m.user, m.err
}

func (m *mockUserDAO) FetchByID(ctx context.Context, id string) (*model.UserEntity, error) {
	return m.user, m.err
}

func (m *mockUserDAO) FetchAll(ctx context.Context) ([]*model.UserEntity, error) {
	return m.users, m.err
}

func (m *mockUserDAO) Update(ctx context.Context, id string, user *model.User) (*model.UserEntity, error) {
	return m.user, m.err
}

func (m *mockUserDAO) Delete(ctx context.Context, id string) error {
	return m.err
}
