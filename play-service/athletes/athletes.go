package athletes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	url := fmt.Sprintf("%s/%s/%d", as.apiUrl, entityType, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("entity not found")
	}

	switch entityType {
	case "athlete":
		type TmpAthlete struct {
			League string `json:"league"`
		}
		var tmp TmpAthlete

		err = json.NewDecoder(res.Body).Decode(&tmp)
		if err != nil {
			return nil, err
		}
		switch tmp.League {
		case "NFL":
			nflStats := models.NFLStats{}
			nflAthlete := models.NFLAthlete{
				Athlete: models.Athlete{
					Stats: &nflStats,
				},
			}
			err = json.NewDecoder(res.Body).Decode(&nflAthlete)
			if err != nil {
				return nil, err
			}
			return &nflAthlete, nil
		case "NBA":
			nbaStats := models.NBAStats{}
			nbaAthlete := models.NBAAthlete{
				Athlete: models.Athlete{
					Stats: &nbaStats,
				},
			}
			err = json.NewDecoder(res.Body).Decode(&nbaAthlete)
			if err != nil {
				return nil, err
			}
			return &nbaAthlete, nil
		default:
			return nil, errors.New("invalid league")
		}
	case "team":
		return nil, errors.New("invalid league")
	}

	return nil, nil
}

// TODO
func (as *athleteService) ListEntitiesByIDs(ids []int, entityType string) ([]models.Entity, error) {
	url := fmt.Sprintf("%s/%s", as.apiUrl, entityType)
	json, err := json.Marshal(ids)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("entities not found")
	}

	entities := make([]models.Entity, 0)

	/*
	 [ entity, entity ]
	*/

	switch entityType {
	case "athlete":
		type TmpAthlete struct {
			League string `json:"league"`
		}
		var tmp TmpAthlete

		err = json.NewDecoder(res.Body).Decode(&tmp)
		if err != nil {
			return nil, err
		}
		switch tmp.League {
		case "NFL":
			nflStats := models.NFLStats{}
			nflAthlete := models.NFLAthlete{
				Athlete: models.Athlete{
					Stats: &nflStats,
				},
			}
			err = json.NewDecoder(res.Body).Decode(&nflAthlete)
			if err != nil {
				return nil, err
			}
			return &nflAthlete, nil
		case "NBA":
			nbaStats := models.NBAStats{}
			nbaAthlete := models.NBAAthlete{
				Athlete: models.Athlete{
					Stats: &nbaStats,
				},
			}
			err = json.NewDecoder(res.Body).Decode(&nbaAthlete)
			if err != nil {
				return nil, err
			}
			return &nbaAthlete, nil
		default:
			return nil, errors.New("invalid league")
		}
	case "team":
		return nil, errors.New("invalid league")
	}

	return nil, nil

	return nil, nil
}
