package service

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	//Services
}

func (s Service) ValidateEmail(e string) (string, error){
	if !strings.Contains(e, "@"){
		return "",fmt.Errorf("invalid email")
	}
	return e, nil
}

func (s *Service) Signup(ctx context.Context, email, password string) error {
		_, err := s.ValidateEmail(email)
		if err != nil {
			return NewError(ErrBadRequest,err)
		}
		return nil
}