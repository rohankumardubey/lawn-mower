package domain

type CatalogRepository interface {
	Find(id string) (*Mower, error)
	Add(input AddMowerDTO) (*Mower, error)
}
