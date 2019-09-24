package greeter

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer http server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		endpoints.StatusEndpoint,
		decodeStatusRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/hello").Handler(httptransport.NewServer(
		endpoints.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	))

	return r
}

// Add specific Headers to every response
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Context-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
