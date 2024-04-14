package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Msg string `json:"error,omitempty"`
}

func SendErrorResponse(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	if msg != "" {
		w.Header().Set("Content-Type", "application/json")
		payload := ErrorResponse{
			Msg: msg,
		}
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			w.Header().Del("Content-Type")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func SendSuccessResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	payload, err := json.Marshal(data)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(payload); err != nil {
		w.Header().Del("Content-Type")
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}
