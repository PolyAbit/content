package content

import "errors"

var (
	ErrDirectionNotFound = errors.New("direction not found")
	ErrCodeAlreadyUsed = errors.New("code already used")
)
