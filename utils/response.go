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

	result, err := json.Marshal(response)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(result)
	return err
}

func ResponseJSON(w http.ResponseWriter, code int, status string, payload interface{}) error {
	response := response{
		Code:   code,
		Status: status,
		Data:   payload,
	}

	result, err := json.Marshal(response)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(result)
	return err
}
