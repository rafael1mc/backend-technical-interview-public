package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Controller) HandleCreatePick(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UserID     int    `json:"userId"`
		EntityID   int    `json:"entityId"`
		EntityType string `json:"entityType"`
	}

	req := Request{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.s.CreatePick(req.UserID, req.EntityID, req.EntityType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) HandleGetRoster(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["userID"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	roster, err := c.s.GetRoster(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(roster)
}
