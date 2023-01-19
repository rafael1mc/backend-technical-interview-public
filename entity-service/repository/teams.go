package repository

import (
	"entity-service/models"
)

// TODO
func (r *Repository) GetTeamByID(id int) (models.Team, error) {
	var team models.Team
	var err error

	return team, err
}

// TODO
func (r *Repository) ListTeamsByIDs(ids []int) ([]models.Team, error) {
	teams := []models.Team{}
	var err error

	return teams, err
}
