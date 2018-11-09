package handlers

import (
	"github.com/BankEx/ebh/services"
	"github.com/BankEx/ebh/repositories"
)

type HandlersService struct {
	service *services.NodeReader
	repository *repositories.MongoRepository
}

func New (
	service *services.NodeReader,
	repository *repositories.MongoRepository) *HandlersService {
	return &HandlersService{
		service: service,
		repository: repository,
	}
}
