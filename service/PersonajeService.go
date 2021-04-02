package service

import (
	"aprendiendo/domain"
	"aprendiendo/repository"
	"errors"
)

type PersonajeService struct {
	repository *repository.PersonajeRepository
}

func NewPersonajeService(repository *repository.PersonajeRepository) *PersonajeService {
	personajeService := &PersonajeService{}
	personajeService.repository = repository
	return personajeService
}

func (s *PersonajeService) Save(data domain.Personaje) error {
	if data.Nombre == "ctm" {
		return errors.New("Tu personaje no puede tener ese nombre")
	}

	s.repository.Save(data)
	return nil

}
