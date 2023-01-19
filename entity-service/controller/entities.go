package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Controller) GetEntityByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	entityType := vars["entityType"]
	entityIDStr := vars["entityID"]

	entityID, err := strconv.Atoi(entityIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity, err := c.s.GetEntityByID(entityID, entityType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(entity)
}

func (c *Controller) ListEntityByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	entityType := vars["entityType"]

	var entityIds []int

	err := json.NewDecoder(r.Body).Decode(&entityIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entities, err := c.s.ListEntitiesByIDs(entityIds, entityType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(entities)
}
