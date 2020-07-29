package errorx

import "errors"

var (
	ErrNoUser = errors.New("user is not present")
	ErrDeleteUser = errors.New("user has been deleted")
)