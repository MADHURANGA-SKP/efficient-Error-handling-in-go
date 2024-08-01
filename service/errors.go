package service

import "errors"

var (
	ErrBadRequest      = errors.New("bad requst")
	ErrInternalFailure = errors.New("internal faliure")
	ErrNotFound        = errors.New("not found")
)

type Error struct {
	appError error
	svcError error
}

func (e Error) AppError() error {
	return e.appError
}

func (e Error) SvcError() error {
	return e.svcError
}

func NewError(svcErr, appErr error) error {
	return Error{
		appError: appErr,
		svcError: svcErr,
	}
}

func (e Error) Error() string {
	return errors.Join(e.svcError, e.appError).Error()
}
