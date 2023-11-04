package handler

import (
	"github.com/mkamadeus/yals/domain"
)

type yalsHandler struct {
	Service domain.YALSService
}

func NewYALSHandler(service domain.YALSService) yalsHandler {
	return yalsHandler{
		Service: service,
	}
}
