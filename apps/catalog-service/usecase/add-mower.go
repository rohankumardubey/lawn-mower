package usecase

import (
	"jrobic/lawn-mower/catalog-service/domain"
)

func (lm *LMCatalogService) Add(input domain.AddMowerDTO) (*domain.Mower, error) {
	mower, err := lm.repo.Add(input)

	if err != nil {
		return nil, err
	}

	return mower, nil
}
