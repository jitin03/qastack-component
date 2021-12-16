package domain

import (
	"fmt"
	"qastack-components/errs"
	logger "qastack-components/loggers"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ComponentRepositoryDb struct {
	client *sqlx.DB
}

func (d ComponentRepositoryDb) AddComponent(c Component) (*Component, *errs.AppError) {

	sqlInsert := "INSERT INTO component (name, project_id) values ($1, $2) RETURNING id"
	var id string
	err := d.client.QueryRow(sqlInsert, c.Name, c.Project_Id).Scan(&id)

	if err != nil {
		logger.Error("Error while creating new component: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	c.Component_Id = id
	return &c, nil
}

func (d ComponentRepositoryDb) AllComponent(projectKey string ,pageId int) ([]Component, *errs.AppError) {
	var err error
	components := make([]Component, 0)
	logrus.Info(projectKey)
	findAllSql := "select id,name, project_id from component where project_id=$1 LIMIT $2"
	err = d.client.Select(&components, findAllSql,projectKey,pageId)

	if err != nil {
		fmt.Println("Error while querying component table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return components, nil
}

func (d ComponentRepositoryDb) DeleteComponent(id int) *errs.AppError {

	deleteSql := "DELETE FROM component WHERE id = $1"
	res, err := d.client.Exec(deleteSql, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	return nil

}

func (d ComponentRepositoryDb) UpdateComponent(id int, newComponent Component) ( *errs.AppError) {

	updateComponentSql := "UPDATE component SET name = $1 WHERE id = $2"
	res, err := d.client.Exec(updateComponentSql,newComponent.Name,id)
	if err != nil {
		return  errs.NewUnexpectedError("Unexpected error from database")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return  errs.NewUnexpectedError("Unexpected error from database")
	}
	fmt.Println(count)
	return nil
}

func NewComponentRepositoryDb(dbClient *sqlx.DB) ComponentRepositoryDb {
	return ComponentRepositoryDb{dbClient}
}
