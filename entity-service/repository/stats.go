package repository

import (
	"entity-service/models"
	"errors"
)

// TODO
func (r *Repository) GetStatsByEntityID(id int, league string) (models.Stats, error) {
	switch league {
	case "NFL":
		var stat models.NFLStats
		err := r.db.Get(
			&stat,
			`SELECT * 
			FROM nfl_stats
			WHERE entity_id = $1`,
			id,
		)

		return stat, err
	case "NBA":
		var stat models.NBAStats
		err := r.db.Get(
			&stat,
			`SELECT * 
				FROM nba_stats
				WHERE entity_id = $1`,
			id,
		)

		return stat, err
	default:
		return nil, errors.New("invalid league")
	}
}

// TODO
func (r *Repository) ListStatsByEntityIDs(ids []int, league string) ([]models.Stats, error) {
	var stats []models.Stats
	var err error

	return stats, err
}
