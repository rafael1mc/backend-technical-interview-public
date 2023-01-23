package service

import (
	"entity-service/models"
	"errors"
	"strings"
)

// TODO
func (s *Service) GetEntityByID(entityID int, entityType string) (models.Entity, error) {
	var entity models.Entity
	var err error

	switch strings.ToLower(entityType) {
	case "athlete":
		athlete, err := s.r.GetAthleteByID(entityID)
		if err != nil {
			return nil, err
		}
		entity = &athlete
	case "team":
		team, err := s.r.GetTeamByID(entityID)
		if err != nil {
			return nil, err
		}
		entity = &team
	default:
		return nil, errors.New("invalid entity type")
	}
	stats, err := s.r.GetStatsByEntityID(entity.GetID(), entity.GetLeague())
	if err != nil {
		return nil, err
	}
	entity.SetStats(stats)

	return entity, err
}

// TODO
func (s *Service) ListEntitiesByIDs(entityIDs []int, entityType string) ([]models.Entity, error) {
	var entities []models.Entity
	var err error

	return entities, err
}
