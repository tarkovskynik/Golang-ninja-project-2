package domain

import "errors"

var (
	NotFoundUserID = errors.New("user id not found in header")
)