package domain

import "errors"

var (
	NotFoundUserID = errors.New("user id not found in header")
	NotFoundUser = errors.New("user doesn't exist by mail")
)
