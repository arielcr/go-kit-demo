package greeter

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints definition
type Endpoints struct {
	StatusEndpoint endpoint.Endpoint
	HelloEndpoint  endpoint.Endpoint
}

// MakeStatusEndpoint returns the response from our service "status"
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest)
		s, err := srv.Status(ctx)

		if err != nil {
			return statusResponse{s}, err
		}
		return statusResponse{s}, nil
	}
}

// MakeHelloEndpoint returns the response from our service "hello"
func MakeHelloEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(helloRequest)
		s, err := srv.Hello(ctx, req.Name)
		if err != nil {
			return helloResponse{s}, err
		}
		return helloResponse{s}, nil
	}
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

// Hello endpoint mapping
func (e Endpoints) Hello(ctx context.Context, name string) (string, error) {
	req := helloRequest{Name: name}
	resp, err := e.HelloEndpoint(ctx, req)
	if err != nil {
		return "", nil
	}
	helloResp := resp.(helloResponse)
	return helloResp.Greetings, nil
}
