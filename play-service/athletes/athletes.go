package athletes

import (
	"os"
	"play-service/models"
)

type AthleteProvider interface {
	GetEntityByID(id int, entityType string) (models.Entity, error)
	ListEntitiesByIDs(ids []int, entityType string) ([]models.Entity, error)
}

type athleteService struct {
	apiUrl string
}

func New() AthleteProvider {
	return newAthleteService()
}

func newAthleteService() *athleteService {
	apiUrl := os.Getenv("ATHLETE_SERVICE_URL")

	return &athleteService{apiUrl}
}

// TODO
func (as *athleteService) GetEntityByID(id int, entityType string) (models.Entity, error) {
	return nil, nil
}

// TODO
func (as *athleteService) ListEntitiesByIDs(ids []int, entityType string) ([]models.Entity, error) {
	return nil, nil
}
