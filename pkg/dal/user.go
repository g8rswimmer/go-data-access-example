package dal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/g8rswimmer/go-data-access-example/pkg/errorx"
	"github.com/g8rswimmer/go-data-access-example/pkg/model"
	"github.com/google/uuid"
)

const uuidLength = 36

type User struct {
	DB *sql.DB
}

func (u *User) Create(ctx context.Context, user *model.User) (*model.UserEntity, error) {
	if user == nil {
		return nil, errors.New("user can not be nil")
	}

	now := time.Now()

	e := &model.UserEntity{
		Entity: model.Entity{
			ID:        uuid.New().String(),
			CreatedAt: now,
			UpdatedAt: now,
		},
		User: model.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}

	const stmt = `INSERT INTO user (id, first_name, last_name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	if _, err := u.DB.ExecContext(ctx, stmt, e.ID, e.FirstName, e.LastName, e.CreatedAt, e.UpdatedAt); err != nil {
		return nil, fmt.Errorf("user create insert %w", err)
	}
	return e, nil
}

func (u *User) FetchByID(ctx context.Context, id string) (*model.UserEntity, error) {
	if len(id) != uuidLength {
		return nil, fmt.Errorf("user fetch by id length %d", len(id))
	}

	const stmt = `SELECT id, first_name, last_name, created_at, updated_at FROM user WHERE id = ? AND updated_at = NULL`
	row := u.DB.QueryRowContext(ctx, stmt, id)

	e := &model.UserEntity{}
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.CreatedAt, &e.UpdatedAt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, errorx.ErrNoUser
	case err != nil:
		return nil, fmt.Errorf("user fetch query %w", err)
	default:
		return e, nil
	}

}

func (u *User) Update(ctx context.Context, id string, user *model.User) (*model.UserEntity, error) {
	switch {
	case len(id) != uuidLength:
		return nil, fmt.Errorf("user fetch by id length %d", len(id))
	case user == nil:
		return nil, errors.New("user can not be nil")
	default:
	}

	e, err := u.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(user.FirstName) == 0 {
		e.FirstName = user.FirstName
	}
	if len(user.LastName) == 0 {
		e.LastName = user.LastName
	}
	e.UpdatedAt = time.Now()

	const stmt = `UPDATE user SET first_name = ?, last_name = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	if _, err := u.DB.ExecContext(ctx, stmt, e.FirstName, e.LastName, id); err != nil {
		return nil, err
	}
	return e, nil
}

func (u *User) Delete(ctx context.Context, id string) error {
	if len(id) != uuidLength {
		return fmt.Errorf("user fetch by id length %d", len(id))
	}

	_, err := u.FetchByID(ctx, id)
	if err != nil {
		return err
	}

	const stmt = `UPDATE user SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`
	if _, err := u.DB.ExecContext(ctx, stmt, id); err != nil {
		return err
	}
	return nil

}
