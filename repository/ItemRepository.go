package repository

import (
	"aprendiendo/domain"
	"context"

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

func (r *ItemRepository) Getall() []domain.Item {

	items := make([]domain.Item, 1)
	dociter := r.collection.Documents(context.Background())
	defer dociter.Stop()
	for {
		doc, _ := dociter.Next()
		var item domain.Item
		doc.DataTo(&item)
		items = append(items, item)
	}

	return items
}
