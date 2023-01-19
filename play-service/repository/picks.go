package repository

import "play-service/models"

// TODO
func (r *Repository) CreatePick(
	userID int,
	entityID int,
	entityType string,
) error {
	return nil
}

// TODO
func (r *Repository) ListPicksByUserID(userID int) ([]models.Pick, error) {
	return nil, nil
}
