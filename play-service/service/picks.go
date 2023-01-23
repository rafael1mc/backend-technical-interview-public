package service

import "play-service/models"

// TODO
func (s *Service) CreatePick(userID int, entityID int, entityType string) error {
	// before inserting, validate if entityId exists (from the entity-service)
	_, err := s.athletes.GetEntityByID(entityID, entityType)
	if err != nil {
		return err
	}

	// can insert pick into database
	err = s.r.CreatePick(userID, entityID, entityType)
	if err != nil {
		return err
	}

	return nil
}

// TODO
func (s *Service) GetRoster(userID int) (models.Roster, error) {
	picks, err := s.r.ListPicksByUserID(userID)
	if err != nil {
		return nil, err
	}

	ids := make([]int, 0)
	for _, pick := range picks {
		ids = append(ids, pick.EntityID)
	}

	entities, err := s.athletes.ListEntitiesByIDs(ids, picks[0].EntityType)

	for _, pick := range picks {
		for _, entity := range entities {
			if pick.EntityID == entity.GetID() {
				pick.Entity = entity
				pick.Fpts = entity.CalculateFpts()
			}
		}
	}

	return models.Roster(picks), nil
}
