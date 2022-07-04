package usecase

import (
	lm_testing "jrobic/lawn-mower/catalog-service"
	domain "jrobic/lawn-mower/catalog-service/domain"
	"strings"

	"testing"
)

func TestFindMower(t *testing.T) {
	t.Run("catalog: return a mower", func(t *testing.T) {
		want := &domain.Mower{Id: "1", Name: "M-350"}

		repo := &lm_testing.StubCatalogRepository{Mowers: map[string]*domain.Mower{
			"1": want,
		}}
		service := NewCatalogService(repo)

		got, err := service.Find("1")

		lm_testing.AssertNoError(t, err)
		lm_testing.AssertMowerEquals(t, *got, *want)
	})

	t.Run("catalog: return error when mower not found", func(t *testing.T) {
		want := &domain.Mower{Id: "1", Name: "M-350"}
		idNotFound := "2"

		repo := &lm_testing.StubCatalogRepository{Mowers: map[string]*domain.Mower{
			"1": want,
		}}
		service := NewCatalogService(repo)

		_, err := service.Find(idNotFound)

		got := strings.Replace(domain.ErrMowerNotFound, "%v", idNotFound, 1)

		lm_testing.AssertError(t, err, got)
	})
}

func TestAddMower(t *testing.T) {
	t.Run("catalog: add mower", func(t *testing.T) {
		repo := &lm_testing.StubCatalogRepository{Mowers: map[string]*domain.Mower{}}
		service := NewCatalogService(repo)

		newMower := domain.AddMowerDTO{Name: "M-150"}

		got, err := service.Add(newMower)

		lm_testing.AssertNoError(t, err)

		if m := repo.Mowers[got.Id]; m == nil {
			t.Errorf("could not find new added mower")
		}
	})
}
