package controller

import "play-service/service"

type Controller struct {
	s *service.Service
}

func New(s *service.Service) *Controller {
	return &Controller{s}
}
