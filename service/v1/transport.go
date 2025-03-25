package v1

import (
	"context"
	"encoding/json"
	"net/http"
)

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req helloRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
