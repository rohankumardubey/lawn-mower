package domain

import "fmt"

type Mower struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AddMowerDTO struct {
	Name string `json:"name"`
}

func NewMower(name string) *Mower {
	return &Mower{
		Id:   "LM-123",
		Name: name,
	}
}

func (m Mower) String() string {
	return fmt.Sprintf("mower %s (#%v)", m.Name, m.Id)
}
