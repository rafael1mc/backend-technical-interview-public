package repository

import (
	"entity-service/models"
)

// TODO
func (r *Repository) GetStatsByEntityID(id int, league string) (models.Stats, error) {
	var stats models.Stats
	var err error

	return stats, err
}

// TODO
func (r *Repository) ListStatsByEntityIDs(ids []int, league string) ([]models.Stats, error) {
	var stats []models.Stats
	var err error

	return stats, err
}
