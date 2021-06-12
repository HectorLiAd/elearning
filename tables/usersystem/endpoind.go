package usersystem

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type loginUserRequest struct {
	Email    string
	Password string
}

func makeLoginUserEndPoint(s Service) endpoint.Endpoint {
	registerUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginUserRequest)
		result, err := s.LoginUser(&req)
		return result, err
	}
	return registerUserEndPoint
}

func mensaje(s Service) endpoint.Endpoint {
	registerUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{"Message": "Correct connection to API"}, nil
	}
	return registerUserEndPoint
}
