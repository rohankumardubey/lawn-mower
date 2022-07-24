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

type AddMowerInputDTO struct {
	Name string
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
	router.HandleFunc("/mowers/", http.HandlerFunc(s.FindMower))
	router.HandleFunc("/", http.HandlerFunc(s.GetCatalog))

	s.Handler = router

	return s, nil
}

func (serv *CatalogHttpServer) CreateMower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)

	mowerToCreate := &AddMowerInputDTO{}

	err := json.NewDecoder(r.Body).Decode(&mowerToCreate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mower, err := serv.service.Add(domain.AddMowerDTO{Name: mowerToCreate.Name})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(mower)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (serv *CatalogHttpServer) FindMower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)

	id := strings.TrimPrefix(r.URL.Path, "/mowers/")

	mower, err := serv.service.Find(id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = json.NewEncoder(w).Encode(mower)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (serv *CatalogHttpServer) GetCatalog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)

	mowers, err := serv.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(mowers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
