package repository

import (
	"aprendiendo/domain"
	"context"

	"cloud.google.com/go/firestore"
)

type PersonajeRepository struct {
	collection *firestore.CollectionRef
}

func NewPersonajeRepository(c *firestore.Client) *PersonajeRepository {
	repository := &PersonajeRepository{}
	repository.collection = c.Collection("personajes")
	return repository
}

func (r *PersonajeRepository) Save(data domain.Personaje) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return r.collection.Add(context.Background(), data)
}
