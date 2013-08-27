package global

import (
	"errors"
)

var (
	InvalidIdError        = errors.New("E001")
	PermissionDeniedError = errors.New("E002")
	ObjectAlreadyExists   = errors.New("E003")
)
