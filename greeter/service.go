package greeter

import "context"

// Service interface
type Service interface {
	Status(ctx context.Context) (string, error)
	Hello(ctx context.Context, name string) (string, error)
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
