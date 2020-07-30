package dal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/g8rswimmer/go-data-access-example/pkg/errorx"
	"github.com/g8rswimmer/go-data-access-example/pkg/model"
)

const uuidLength = 36

// User handles all of the database actions
type User struct {
	DB           *sql.DB
	GenerateUUID GenerateUUID
}

// Create will insert a user into the database
func (u *User) Create(ctx context.Context, user *model.User) (*model.UserEntity, error) {
	if user == nil {
		return nil, errors.New("user can not be nil")
	}

	now := time.Now()

	e := &model.UserEntity{
		Entity: model.Entity{
			ID:        u.GenerateUUID(),
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

// FetchByID returns an entity by the id
func (u *User) FetchByID(ctx context.Context, id string) (*model.UserEntity, error) {
	if len(id) != uuidLength {
		return nil, fmt.Errorf("user fetch by id length %d", len(id))
	}

	const stmt = `SELECT id, first_name, last_name, created_at, updated_at, deleted_at FROM user WHERE id = ?`
	row := u.DB.QueryRowContext(ctx, stmt, id)

	e := &model.UserEntity{}
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.CreatedAt, &e.UpdatedAt, &e.DeletedAt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, errorx.ErrNoUser
	case err != nil:
		return nil, fmt.Errorf("user fetch query %w", err)
	case e.DeletedAt.Valid:
		return nil, errorx.ErrDeleteUser
	default:
		return e, nil
	}

}

// FetchAll returns all entities
func (u *User) FetchAll(ctx context.Context) ([]*model.UserEntity, error) {

	const stmt = `SELECT id, first_name, last_name, created_at, updated_at, deleted_at FROM user`
	rows, err := u.DB.QueryContext(ctx, stmt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, errorx.ErrNoUser
	case err != nil:
		return nil, fmt.Errorf("user fetch query %w", err)
	default:
	}
	defer rows.Close()

	entities := []*model.UserEntity{}
	for rows.Next() {
		e := &model.UserEntity{}
		deletedAt := sql.NullTime{}
		if err := rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.CreatedAt, &e.UpdatedAt, &deletedAt); err != nil {
			return nil, fmt.Errorf("user row scan error %w", err)
		}
		if deletedAt.Valid == false {
			entities = append(entities, e)
		}
	}

	return entities, nil
}

// Update will update an entity with new information
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

	if len(user.FirstName) > 0 {
		e.FirstName = user.FirstName
	}
	if len(user.LastName) > 0 {
		e.LastName = user.LastName
	}
	e.UpdatedAt = time.Now()

	const stmt = `UPDATE user SET first_name = ?, last_name = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	if _, err := u.DB.ExecContext(ctx, stmt, e.FirstName, e.LastName, id); err != nil {
		return nil, err
	}
	return e, nil
}

// Delete will soft delete an entity
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
