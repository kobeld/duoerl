package global

import (
	"errors"
)

var (
	InvalidIdError        = errors.New("Invalid Id")
	PermissionDeniedError = errors.New("Permission denied")
)
