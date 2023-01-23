package repository

import "play-service/models"

// TODO
func (r *Repository) CreatePick(
	userID int,
	entityID int,
	entityType string,
) error {
	sql := `
		INSERT INTO picks VALUES (DEFAULT, $1, $2, $3)
	`
	_, err := r.db.Exec(sql, userID, entityID, entityType)
	return err
}

// TODO
func (r *Repository) ListPicksByUserID(userID int) ([]models.Pick, error) {
	picks := []models.Pick{}
	err := r.db.Select(
		&picks,
		`SELECT *
		FROM picks
		WHERE user_id = $1`,
		userID,
	)

	return picks, err
}
