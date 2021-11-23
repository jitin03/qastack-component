package domain

import (
	"qastack-components/dto"
	"qastack-components/errs"
)

type Component struct {
	Component_Id int    `db:"id"`
	Name         string `db:"name"`
	Project_Id   int    `db:"project_id"`
}

func (c Component) ToDto() dto.AddComponentResponse {
	return dto.AddComponentResponse{
		Component_Id: c.Component_Id,
		Name:       c.Name,
		Project_id: c.Project_Id,
	}
}

func (c Component) ToAddComponentResponseDto() *dto.AddComponentResponse {
	return &dto.AddComponentResponse{c.Component_Id, c.Name, c.Project_Id}
}

type ComponentRepository interface {
	AddComponent(component Component) (*Component, *errs.AppError)
	AllComponent() ([]Component, *errs.AppError)
	DeleteComponent(id int)(*errs.AppError)
	UpdateComponent(id int,newComponent Component)( *errs.AppError)
}
