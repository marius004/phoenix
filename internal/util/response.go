package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func DataResponse(w http.ResponseWriter, statusCode int, retData interface{}, logger *log.Logger) {
	if err, ok := retData.(error); ok {
		retData = err.Error()
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(retData)

	if err != nil {
		logger.Printf("[ERROR] Could not return json data %v %v\n", err, retData)
	}
}

func EmptyResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, retData interface{}, logger *log.Logger) {
	if err, ok := retData.(error); ok {
		retData = err.Error()
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(struct {
		Message interface{} `json:"message"`
	}{
		Message: retData,
	})

	if err != nil {
		logger.Printf("[ERROR] Could not return json data %v %v\n", err, retData)
	}
}
