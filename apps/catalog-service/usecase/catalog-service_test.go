package usecase

import (
	lm_testing "jrobic/lawn-mower/catalog-service"
	domain "jrobic/lawn-mower/catalog-service/domain"
	"reflect"
	"strings"

	"testing"
)

func TestGetMower(t *testing.T) {
	t.Run("catalog: return a mower", func(t *testing.T) {
		want := &domain.Mower{Id: "1", Name: "M-350"}

		wantedCatalog := []*domain.Mower{
			want,
		}
		repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}

		service := NewCatalogService(repo)

		got, err := service.GetMower("1")

		lm_testing.AssertNoError(t, err)
		lm_testing.AssertMowerEquals(t, *got, *want)
	})

	t.Run("catalog: return error when mower not found", func(t *testing.T) {
		want := &domain.Mower{Id: "1", Name: "M-350"}
		wantedCatalog := []*domain.Mower{
			want,
		}

		repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}

		idNotFound := "2"

		service := NewCatalogService(repo)

		_, err := service.GetMower(idNotFound)

		got := strings.Replace(domain.ErrMowerNotFound, "%v", idNotFound, 1)

		lm_testing.AssertError(t, err, got)
	})
}

func TestCreateMower(t *testing.T) {
	t.Run("catalog: create new mower", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{}
		repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}
		service := NewCatalogService(repo)

		newMower := domain.CreateMowerDTO{Name: "M-150"}

		insertedMower, err := service.CreateMower(newMower)

		lm_testing.AssertNoError(t, err)

		got, _ := service.GetMower(insertedMower.Id)

		if got == nil {
			t.Errorf("could not find new created mower")
		}
	})
}

func TestUpdateMower(t *testing.T) {
	t.Run("catalog: update mower", func(t *testing.T) {
		wantedMower := domain.Mower{Id: "1", Name: "M-90"}

		wantedUpdatedMower := wantedMower
		wantedUpdatedMower.Name = "M-150"

		wantedCatalog := []*domain.Mower{
			&wantedMower,
		}

		repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}
		service := NewCatalogService(repo)

		updateMower := domain.UpdateMowerDTO{Name: wantedUpdatedMower.Name}

		_, err := service.UpdateMower(wantedCatalog[0].Id, updateMower)

		lm_testing.AssertNoError(t, err)

		got, _ := service.GetMower(wantedCatalog[0].Id)

		if got == nil {
			t.Errorf("could not find updated mower")
			return
		}

		lm_testing.AssertMowerEquals(t, wantedUpdatedMower, *got)
	})

	t.Run("catalog: update mower with empty 'Name'", func(t *testing.T) {
		wantedMower := domain.Mower{Id: "1", Name: "M-90"}

		wantedCatalog := []*domain.Mower{
			&wantedMower,
		}

		repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}
		service := NewCatalogService(repo)

		updateMower := domain.UpdateMowerDTO{}

		_, err := service.UpdateMower(wantedCatalog[0].Id, updateMower)

		lm_testing.AssertNoError(t, err)

		got, _ := service.GetMower(wantedCatalog[0].Id)

		if got == nil {
			t.Errorf("could not find updated mower")
			return
		}

		lm_testing.AssertMowerEquals(t, wantedMower, *got)
	})
}

func TestGetAvailableMowers(t *testing.T) {
	t.Run("catalog: find all mowers", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{
			{Id: "1", Name: "M-90"},
			{Id: "2", Name: "M-150"},
			{Id: "3", Name: "M-480"},
		}
		repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}

		wantedMowers := []*domain.Mower{
			{Id: "1", Name: "M-90"},
			{Id: "2", Name: "M-150"},
			{Id: "3", Name: "M-480"},
		}

		service := NewCatalogService(repo)

		got, err := service.GetAvailableMowers()

		lm_testing.AssertNoError(t, err)

		if len(got) != len(wantedMowers) {
			t.Errorf("got %v want %v", len(got), len(wantedMowers))
		}

		if !reflect.DeepEqual(got, wantedMowers) {
			t.Errorf("got %v want %v", got, wantedMowers)
		}
	})
}
