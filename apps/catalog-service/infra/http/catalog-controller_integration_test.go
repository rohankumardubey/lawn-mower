package http_controller

import (
	"fmt"
	lm_testing "jrobic/lawn-mower/catalog-service"
	"jrobic/lawn-mower/catalog-service/domain"
	"jrobic/lawn-mower/catalog-service/infra/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMowersAndRetrievingThemes(t *testing.T) {
	repo := repository.NewInMemoryRepo([]*domain.Mower{})
	server, _ := NewCatalogHttpServer(repo)

	wantedMowers := []*domain.Mower{
		{Id: "1", Name: "M-90"},
		{Id: "2", Name: "M-150"},
		{Id: "3", Name: "M-480"},
	}

	for _, wantedMower := range wantedMowers {
		server.ServeHTTP(httptest.NewRecorder(), NewCreateMowerRequest(&CreateMowerInputDTO{Name: wantedMower.Name}))
	}

	for _, wantedMower := range wantedMowers {
		testCaseName := "get " + wantedMower.Name + " mower"

		t.Run(testCaseName, func(t *testing.T) {
			response := httptest.NewRecorder()

			server.ServeHTTP(response, NewGetMowerRequest(wantedMower.Id))
			got := lm_testing.GetMowerFromResponse(t, response.Body)

			lm_testing.AssertStatus(t, response.Code, http.StatusOK)
			lm_testing.AssertMowerEquals(t, got, *wantedMower)
		})
	}

	t.Run("get list mowers", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetCatalogRequest())

		got := lm_testing.GetCatalogFromResponse(t, response.Body)

		lm_testing.AssertCatalogEquals(t, got, wantedMowers)
		lm_testing.AssertStatus(t, response.Code, http.StatusOK)
		lm_testing.AssertContentType(t, response, JsonContentType)
	})
}

func TestUpdateMowersAndRetrievingThemes(t *testing.T) {
	wantedMowers := []*domain.Mower{
		{Id: "1", Name: "M-90"},
		{Id: "2", Name: "M-150"},
		{Id: "3", Name: "M-480"},
	}

	wantedUpdatedMowers := []*domain.Mower{
		{Id: "1", Name: "M-90"},
		{Id: "2", Name: "M-150"},
		{Id: "3", Name: "M-390"},
	}

	repo := repository.NewInMemoryRepo(wantedMowers)
	server, _ := NewCatalogHttpServer(repo)

	for _, wantedMower := range wantedUpdatedMowers {
		server.ServeHTTP(httptest.NewRecorder(), NewUpdateMowerRequest(wantedMower.Id, &UpdateMowerInputDTO{Name: wantedMower.Name}))
	}

	for _, wantedMower := range wantedUpdatedMowers {
		testCaseName := "patch " + wantedMower.Name + " mower"

		t.Run(testCaseName, func(t *testing.T) {
			response := httptest.NewRecorder()

			server.ServeHTTP(response, NewGetMowerRequest(wantedMower.Id))
			got := lm_testing.GetMowerFromResponse(t, response.Body)

			lm_testing.AssertStatus(t, response.Code, http.StatusOK)
			lm_testing.AssertMowerEquals(t, got, *wantedMower)
		})
	}
}

func BenchmarkCreateMower(b *testing.B) {
	repo := repository.NewInMemoryRepo([]*domain.Mower{})
	server, _ := NewCatalogHttpServer(repo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := fmt.Sprintf("M-90 %v", i)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewCreateMowerRequest(&CreateMowerInputDTO{Name: name}))
	}
}

func BenchmarkGetMower(b *testing.B) {
	mowers := []*domain.Mower{}

	for i := 0; i < 200; i++ {
		id := fmt.Sprintf("%d", i)
		name := fmt.Sprintf("M-%d-50", i)

		mowers = append(mowers, &domain.Mower{Id: id, Name: name})
	}

	repo := repository.NewInMemoryRepo(mowers)
	server, _ := NewCatalogHttpServer(repo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetMowerRequest("120"))
	}
}

func BenchmarkGetCatalog(b *testing.B) {
	mowers := []*domain.Mower{}

	for i := 0; i < 200; i++ {
		id := fmt.Sprintf("%d", i)
		name := fmt.Sprintf("M-%d-50", i)

		mowers = append(mowers, &domain.Mower{Id: id, Name: name})
	}

	repo := repository.NewInMemoryRepo(mowers)
	server, _ := NewCatalogHttpServer(repo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetCatalogRequest())
	}
}
