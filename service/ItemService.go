package service

import (
	"aprendiendo/domain"
	"aprendiendo/repository"
	"errors"
)

type ItemService struct{ repository *repository.ItemRepository }

func NewItemService(repository *repository.ItemRepository) *ItemService {
	itemService := &ItemService{}
	itemService.repository = repository
	return itemService
}
func (svc *ItemService) Save(data domain.Item) error {
	if data.Nombre == "animal" {
		return errors.New("nombre invalido")
	}

	svc.repository.Save(data)
	return nil

}
func (svc *ItemService) Getall() []domain.Item {
	return	svc.Getall()

}