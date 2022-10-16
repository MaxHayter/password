package controller

import (
	"github.com/MaxHayter/password/internal/service"
	api "github.com/MaxHayter/password/password"
)

type Controller struct {
	service *service.Service
	api.UnimplementedPasswordServiceServer
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}
