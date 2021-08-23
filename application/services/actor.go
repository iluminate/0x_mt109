package services

import (
	"0x_mt109/application/models"
	"0x_mt109/application/repositories"
)

type ActorService struct {
	repo repositories.IActorRepository
}

func NewActorService(repo repositories.IActorRepository) *ActorService {
	return &ActorService{repo: repo}
}

func (service *ActorService) FindAll() (*[]models.Actor, error) {
	filters := make(map[string]string)
	actors, err := service.repo.Find(filters)
	if err != nil {
		return &actors, err
	}
	return &actors, nil
}

func (service *ActorService) Update(request models.Actor) error {
	err := service.repo.Update(request)
	if err != nil {
		return err
	}
	return nil
}

func (service *ActorService) Create(request models.Actor) error {
	err := service.repo.Create(request)
	if err != nil {
		return err
	}
	return nil
}

func (service *ActorService) Delete(id int) error {
	err := service.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
