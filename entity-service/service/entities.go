package service

import (
	"entity-service/models"
)

// TODO
func (s *Service) GetEntityByID(entityID int, entityType string) (models.Entity, error) {
	var entity models.Entity
	var err error

	return entity, err
}

// TODO
func (s *Service) ListEntitiesByIDs(entityIDs []int, entityType string) ([]models.Entity, error) {
	var entities []models.Entity
	var err error

	return entities, err
}
