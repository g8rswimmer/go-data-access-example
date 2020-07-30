package errorx

import "errors"

var (
	// ErrNoUser when no user entity is found
	ErrNoUser = errors.New("user is not present")
	// ErrDeleteUser when the user has been deleted
	ErrDeleteUser = errors.New("user has been deleted")
)