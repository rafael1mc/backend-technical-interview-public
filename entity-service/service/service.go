package service

import (
	"entity-service/repository"
)

type Service struct {
	r *repository.Repository
}

func New(r *repository.Repository) *Service {
	return &Service{r}
}
