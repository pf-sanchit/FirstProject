package v1

import "context"

type Service interface {
	Hello(ctx context.Context, name string) (message string, errorObj error)
}

type mainService struct{}

func NewService() Service {
	return mainService{}
}

func (mainService) Hello(ctx context.Context, name string) (message string, errorObj error) {
	message = "Hello " + name
	return
}
