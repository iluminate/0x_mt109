package repositories

import "0x_mt109/application/models"

type IActorRepository interface {
	Find(filters map[string]string) ([]models.Actor, error)
	Update(request models.Actor) error
	Create(request models.Actor) error
	Delete(id int) error
}