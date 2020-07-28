package errorx

import "errors"

var (
	ErrNoUser = errors.New("user is not present")
)