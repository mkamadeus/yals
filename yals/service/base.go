package service

import (
	"github.com/mkamadeus/yals/domain"
)

type yalsService struct {
	Repository domain.YALSRepository
}

func NewYALSService(repository domain.YALSRepository) yalsService {
	return yalsService{
		Repository: repository,
	}
}
