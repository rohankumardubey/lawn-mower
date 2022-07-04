package usecase

import (
	"jrobic/lawn-mower/catalog-service/domain"
)

type CatalogService interface {
	Add(input domain.AddMowerDTO) (*domain.Mower, error)
	Find(id string) (*domain.Mower, error)
	FindAll() ([]*domain.Mower, error)
}

type LMCatalogService struct {
	repo domain.CatalogRepository
}

func NewCatalogService(repo domain.CatalogRepository) *LMCatalogService {
	return &LMCatalogService{
		repo: repo,
	}
}
