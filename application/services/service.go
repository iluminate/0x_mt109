package services

import "0x_mt109/application/models"

type IActorService interface {
	FindAll() (*[]models.Actor, error)
	Update(request models.Actor) error
	Create(request models.Actor) error
	Delete(id int) error
}