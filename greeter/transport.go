package greeter

import (
	"context"
	"encoding/json"
	"net/http"
)

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Greetings string `json:"greetings"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}

type complexRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	ZipCode string `json:"zipcode"`
}

type complexResponse struct {
	Data string         `json:"data"`
	Meta complexRequest `json:"meta"`
}

func decodeHelloRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req helloRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeComplexRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req complexRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
