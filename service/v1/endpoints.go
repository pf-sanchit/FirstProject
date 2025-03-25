package v1

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	HelloEndpoint endpoint.Endpoint
}

func MakeHelloEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(helloRequest)
		msg, err := srv.Hello(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return helloResponse{msg, ""}, nil
	}
}

func (e Endpoints) Hello(ctx context.Context, name string) (string, error) {
	req := helloRequest{Name: name}
	msg, err := e.HelloEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	helloResponse := msg.(helloResponse)
	if helloResponse.Err != "" {
		return "", errors.New(helloResponse.Err)
	}
	return helloResponse.Message, nil
}
