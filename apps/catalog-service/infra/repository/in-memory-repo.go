package repository

import (
	"fmt"
	"jrobic/lawn-mower/catalog-service/domain"
	"sync"
)

type InMemoryRepo struct {
	Mowers []*domain.Mower
	lock   sync.RWMutex
}

func NewInMemoryRepo(initialMowers []*domain.Mower) *InMemoryRepo {
	mowers := []*domain.Mower{}

	for _, mower := range initialMowers {
		mowers = append(mowers, mower)
	}

	return &InMemoryRepo{Mowers: mowers}
}

func (r *InMemoryRepo) Add(input domain.AddMowerDTO) (*domain.Mower, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	id := fmt.Sprint(len(r.Mowers) + 1)

	mower := &domain.Mower{
		Id:   id,
		Name: input.Name,
	}

	r.Mowers = append(r.Mowers, mower)

	return mower, nil
}

func (r *InMemoryRepo) Find(id string) (mower *domain.Mower, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i, mower := range r.Mowers {
		if mower.Id == id {
			return r.Mowers[i], nil
		}
	}
	return nil, nil
}

func (r *InMemoryRepo) FindAll() (mowers []*domain.Mower, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.Mowers, nil
}
