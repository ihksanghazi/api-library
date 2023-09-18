package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseError(w http.ResponseWriter, code int, errorMessage interface{}) error {
	response := map[string]interface{}{
		"message": errorMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	return err
}

func ResponseJSON(w http.ResponseWriter, code int, status string, payload interface{}) error {
	response := response{
		Code:   code,
		Status: status,
		Data:   payload,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	return err
}
