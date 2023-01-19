package service

import (
	"play-service/athletes"
	"play-service/repository"
)

type Service struct {
	r        *repository.Repository
	athletes athletes.AthleteProvider
}

func New(r *repository.Repository, athletes athletes.AthleteProvider) *Service {
	return &Service{r, athletes}
}
