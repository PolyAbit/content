package models

import "errors"

var (
	ErrDirectionExists = errors.New("direction with same code already exists")
)
