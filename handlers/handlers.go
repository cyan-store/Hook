package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Webhook not found!",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(json)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(ErrorResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method not allowed!",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(json)
}
