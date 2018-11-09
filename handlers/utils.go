package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type contextKey string

const (
	authDataKey = contextKey("AuthData")
)

func jsonErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	jsonResponse(w, map[string]string{"message": err.Error()}, statusCode)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, string(jsonData))
}
