package greeter

import (
	"context"
	"fmt"
)

// Service interface
type Service interface {
	Status(ctx context.Context) (string, error)
	Hello(ctx context.Context, name string) (string, error)
	Complex(ctx context.Context, request complexRequest) (string, complexRequest, error)
}

type helloService struct{}

// NewService returns a new helloService
func NewService() Service {
	return helloService{}
}

func (helloService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

func (helloService) Hello(ctx context.Context, name string) (string, error) {
	return "Hello " + name + "!", nil
}

func (helloService) Complex(ctx context.Context, request complexRequest) (string, complexRequest, error) {
	//return "Hello " + request.Name + "! Happy to know you live in ", nil
	greeting := fmt.Sprintf(
		"Hello %s! Happy to know you live in %s. I will call you to your number %s and send you a letter to %s",
		request.Name, request.Address, request.Phone, request.ZipCode,
	)
	return greeting, request, nil
}
