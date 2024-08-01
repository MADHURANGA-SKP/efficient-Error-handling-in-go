package main

import (
	"context"
	"errors"
	httperror "errr/httpError"
	"errr/service"
)

func main(){
	ctx := context.Background()
	s := service.Service{}

	err := s.Signup(ctx, "email@email.com", "userPassword")
	if err != nil {
		if errors.Is(err, service.ErrBadRequest){
			return httperror.FromError(err)
		}
	}
}