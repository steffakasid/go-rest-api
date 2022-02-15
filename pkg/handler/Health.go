package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Health struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	health := Health{Status: "Everything ok"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(health)
}
