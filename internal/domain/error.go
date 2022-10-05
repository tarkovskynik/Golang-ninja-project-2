package domain

import "errors"

var (
	NotFoundUser = errors.New("user doesn't exist by mail")
)
