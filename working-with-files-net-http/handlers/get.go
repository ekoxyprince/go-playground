package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		resp := Response{
			Success:    false,
			Message:    "Method not Allowed",
			StatusCode: http.StatusMethodNotAllowed,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := Response{
		Success:    true,
		Message:    "Request successful",
		StatusCode: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}
