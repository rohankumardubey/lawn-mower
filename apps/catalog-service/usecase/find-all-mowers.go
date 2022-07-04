package usecase

import (
	"jrobic/lawn-mower/catalog-service/domain"
)

func (lm *LMCatalogService) FindAll() ([]*domain.Mower, error) {
	return lm.repo.FindAll()
}
