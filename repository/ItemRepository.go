package repository

import (
	"aprendiendo/domain"
	"context"
	"errors"

	"cloud.google.com/go/firestore"
)

type ItemRepository struct{ collection *firestore.CollectionRef }

func NewItemRepository(c *firestore.Client) *ItemRepository {
	repository := &ItemRepository{}
	repository.collection = c.Collection("Item")
	return repository
}

func (r *ItemRepository) Save(data domain.Item) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	return r.collection.Add(context.Background(), data)
}

func (r *ItemRepository) Getall() ([]domain.Item, error) {

	documentIterator := r.collection.Documents(context.Background())

	snapshots, err := documentIterator.GetAll()

	if err != nil {
		return nil, errors.New("Error retrieving documentRefs")
	}

	var items []domain.Item

	for i := 0; i < len(snapshots); i++ {
		var item domain.Item
		snapshots[i].DataTo(&item)
		items = append(items, item)
	}

	return items, nil
}
