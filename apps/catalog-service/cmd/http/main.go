package main

import (
	"jrobic/lawn-mower/catalog-service/domain"
	http_controller "jrobic/lawn-mower/catalog-service/infra/http"
	"jrobic/lawn-mower/catalog-service/infra/repository"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewInMemoryRepo([]*domain.Mower{})

	server, err := http_controller.NewCatalogHttpServer(repo)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	log.Println("Listen on port 5001")

	if err := http.ListenAndServe(":5001", server); err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}

}
