package controller

import "Sinekod/service"


type Controller struct{
	Serivce *service.Service
}

func NewController (service *service.Service) *Controller{
	return &Controller{
		Serivce: service,
	}
}