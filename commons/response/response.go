package response

import (
	"encoding/json"
	"net/http"
)

type JsonEncoderResponse[T any] struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
}

func JsonEncoder[T any](w http.ResponseWriter, code int, message string, success bool, payload T) {
	responseSuccess := &JsonEncoderResponse[T]{
		Message: message,
		Success: success,
		Data:    payload,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseSuccess)
}
