package repository

import (
	"entity-service/models"

	"github.com/lib/pq"
)

func (r *Repository) GetAthleteByID(id int) (models.Athlete, error) {
	var athlete models.Athlete
	err := r.db.Get(
		&athlete,
		`SELECT *
		FROM athletes
		WHERE id = $1`,
		id,
	)

	return athlete, err
}

func (r *Repository) ListAthletesByIDs(ids []int) ([]models.Athlete, error) {
	athletes := []models.Athlete{}
	err := r.db.Select(
		&athletes,
		`SELECT *
		FROM athletes
		WHERE id = ANY($1)`,
		pq.Array(ids),
	)

	return athletes, err
}
