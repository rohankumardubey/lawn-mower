package http_controller

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	lm_testing "jrobic/lawn-mower/catalog-service"
	"jrobic/lawn-mower/catalog-service/domain"
)

func TestCreateMowerCtrl(t *testing.T) {

	t.Run("AddMowerCtrl return accepted on POST", func(t *testing.T) {
		wantedCatalog := []*domain.Mower{
			{Id: "1", Name: "M-90"},
			{Id: "2", Name: "M-150"},
			{Id: "3", Name: "M-480"},
		}
		repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}
		server, _ := NewCatalogHttpServer(repo)

		mower := &AddMowerInputDTO{Name: "M-600"}
		wantedMower := domain.Mower{Name: "M-600", Id: "4"}

		request := NewPostAddMowerRequest(mower)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := lm_testing.GetMowerFromResponse(t, response.Body)

		lm_testing.AssertStatus(t, response.Code, http.StatusAccepted)
		lm_testing.AssertContentType(t, response, JsonContentType)

		lm_testing.AssertMowerEquals(t, got, wantedMower)
	})
}

func TestFindMowerCtrl(t *testing.T) {
	wantedCatalog := []*domain.Mower{
		{Id: "1", Name: "M-90"},
		{Id: "2", Name: "M-150"},
		{Id: "3", Name: "M-480"},
	}

	repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}
	server, _ := NewCatalogHttpServer(repo)

	t.Run("FindMowerCtrl return M-350 mower", func(t *testing.T) {
		request := NewFindAddMowerRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := lm_testing.GetMowerFromResponse(t, response.Body)
		wantedMower := domain.Mower{Id: "1", Name: "M-90"}

		lm_testing.AssertStatus(t, response.Code, http.StatusOK)
		lm_testing.AssertContentType(t, response, JsonContentType)

		lm_testing.AssertMowerEquals(t, got, wantedMower)
	})

	t.Run("FindMowerCtrl return M-150 mower", func(t *testing.T) {
		request := NewFindAddMowerRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := lm_testing.GetMowerFromResponse(t, response.Body)
		wantedMower := domain.Mower{Id: "2", Name: "M-150"}

		lm_testing.AssertStatus(t, response.Code, http.StatusOK)
		lm_testing.AssertContentType(t, response, JsonContentType)

		lm_testing.AssertMowerEquals(t, got, wantedMower)
	})

	t.Run("FindMowerCtrl return 404 on missing mower", func(t *testing.T) {
		request := NewFindAddMowerRequest("6")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		lm_testing.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestGetCatalogCtrl(t *testing.T) {
	wantedCatalog := []*domain.Mower{
		{Id: "1", Name: "M-90"},
		{Id: "2", Name: "M-150"},
		{Id: "3", Name: "M-480"},
	}

	repo := &lm_testing.StubCatalogRepository{Mowers: wantedCatalog}
	server, _ := NewCatalogHttpServer(repo)

	t.Run("GetCatalogCtrl return list of mowers", func(t *testing.T) {
		request := NewGetCatalogRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := lm_testing.GetCatalogFromResponse(t, response.Body)

		lm_testing.AssertCatalogEquals(t, got, wantedCatalog)
		lm_testing.AssertStatus(t, response.Code, http.StatusOK)
		lm_testing.AssertContentType(t, response, JsonContentType)
	})
}

func NewPostAddMowerRequest(body interface{}) *http.Request {
	jsonBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/mowers", bytes.NewReader(jsonBytes))
	return req
}

func NewFindAddMowerRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/mowers/"+id, nil)
	return req
}

func NewGetCatalogRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	return req
}
