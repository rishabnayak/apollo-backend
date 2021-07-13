package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *api) healthCheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]string{
		"status": "available",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}