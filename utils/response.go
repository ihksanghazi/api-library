package utils

import (
	"encoding/json"
	"net/http"
)

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
