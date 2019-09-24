package greeter

import (
	"context"
	"testing"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()

	s, err := srv.Status(ctx)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	ok := s == "ok"

	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestHello(t *testing.T) {
	srv, ctx := setup()

	s, err := srv.Hello(ctx, "You")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing hello
	ok := s == "Hello You!"

	if !ok {
		t.Errorf("expected service to be Hello You!")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
