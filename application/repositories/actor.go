package repositories

import (
	"0x_mt109/application/models"
	"0x_mt109/helpers/database"
)

type ActorRepository struct{
	database *database.MysqldbHelper
}

func NewActorRepository(helper *database.MysqldbHelper) *ActorRepository {
	return &ActorRepository{database: helper}
}

func (repo *ActorRepository) Find(filters map[string]string) ([]models.Actor, error) {
	var actors []models.Actor
	err := repo.database.OpenConnection()
	if err != nil {
		return actors, err
	}
	stmt, err := repo.database.GetConnection().Prepare(`select id, name from actors`)
	defer stmt.Close()
	if err != nil {
		return actors, err
	}
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		return actors, err
	}
	for rows.Next() {
		var actor models.Actor
		err = rows.Scan(&actor.Id, &actor.Name)
		actors = append(actors, actor)
		if err != nil {
			return actors, err
		}
	}
	return actors, nil
}

func (repo *ActorRepository) Update(request models.Actor) error {
	err := repo.database.OpenConnection()
	if err != nil {
		return err
	}
	stmt, err := repo.database.GetConnection().Prepare(`update actors set name = ? where id = ?`)
	defer stmt.Close()
	_, err = stmt.Exec(request.Name, request.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActorRepository) Create(request models.Actor) error {
	err := repo.database.OpenConnection()
	if err != nil {
		return err
	}
	stmt, err := repo.database.GetConnection().Prepare(`insert into actors (name) values (?)`)
	defer stmt.Close()
	_, err = stmt.Exec(request.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActorRepository) Delete(id int) error {
	err := repo.database.OpenConnection()
	if err != nil {
		return err
	}
	stmt, err := repo.database.GetConnection().Prepare(`delete from actors where id = ?`)
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
