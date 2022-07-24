package http_controller

import (
	"encoding/json"
	"jrobic/lawn-mower/catalog-service/domain"
	"jrobic/lawn-mower/catalog-service/usecase"
	"net/http"
	"strings"
)

var (
	JsonContentType = "application/json"
)

type CreateMowerInputDTO struct {
	Name string `json:"name"`
}

type UpdateMowerInputDTO struct {
	Name string `json:"name,omitempty"`
}

type CatalogHttpServer struct {
	http.Handler
	repo    domain.CatalogRepository
	service usecase.CatalogService
}

func NewCatalogHttpServer(repo domain.CatalogRepository) (*CatalogHttpServer, error) {
	s := new(CatalogHttpServer)

	s.repo = repo
	s.service = usecase.NewCatalogService(repo)

	router := http.NewServeMux()
	router.HandleFunc("/mowers", http.HandlerFunc(s.CreateMower))
	router.HandleFunc("/mowers/", http.HandlerFunc(s.GetOrUpdateMower))
	router.HandleFunc("/", http.HandlerFunc(s.GetCatalog))

	s.Handler = router

	return s, nil
}

func (serv *CatalogHttpServer) CreateMower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)

	mowerToCreate := &CreateMowerInputDTO{}

	err := json.NewDecoder(r.Body).Decode(&mowerToCreate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mower, err := serv.service.CreateMower(domain.CreateMowerDTO{Name: mowerToCreate.Name})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(mower)
}

func (serv *CatalogHttpServer) GetOrUpdateMower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)

	id := strings.TrimPrefix(r.URL.Path, "/mowers/")

	mower := &domain.Mower{}
	var err error

	switch r.Method {
	case http.MethodPatch:
		{

			mowerToUpdate := &UpdateMowerInputDTO{}
			err = json.NewDecoder(r.Body).Decode(&mowerToUpdate)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			mower, err = serv.service.UpdateMower(id, domain.UpdateMowerDTO{
				Name: mowerToUpdate.Name,
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	case http.MethodGet:
		{
			mower, err = serv.service.GetMower(id)

			if err != nil {
				http.NotFound(w, r)
				return
			}
		}
	}

	json.NewEncoder(w).Encode(mower)
}

func (serv *CatalogHttpServer) GetCatalog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)

	mowers, err := serv.service.GetAvailableMowers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(mowers)
}
