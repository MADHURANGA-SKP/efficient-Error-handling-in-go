package httperror

import (
	"errors"
	"errr/service"
	"net/http"
)

type ApiError struct {
	Status int
	Message string
}

func FromError(err error) ApiError {
	var apiError ApiError
	var svcError service.Error
	if errors.As(err, &svcError){
		apiError.Message = svcError.AppError().Error()
		svcErr := svcError.SvcError()
		switch svcErr {
			case service.ErrBadRequest:
				apiError.Status = http.StatusBadRequest
			case service.ErrInternalFailure:
				apiError.Status = http.StatusInternalServerError
			case service.ErrNotFound:
				apiError.Status = http.StatusNotFound
		}
	}
	return apiError
}