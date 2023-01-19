package repository

import "entity-service/database"

type Repository struct {
	db *database.Database
}

func New(d *database.Database) *Repository {
	return &Repository{d}
}
