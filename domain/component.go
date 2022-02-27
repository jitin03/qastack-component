package domain

import (
	"qastack-components/dto"
	"qastack-components/errs"
)

type Component struct {
	Component_Id string `db:"id"`
	Name         string `db:"name"`
	Project_Id   string `db:"project_id"`
	CreateDate   string `db:"create_date"`
	UpdateDate   string `db:"update_date"`
}

func (c Component) ToDto() dto.AddComponentResponse {
	return dto.AddComponentResponse{
		Component_Id: c.Component_Id,
		Name:         c.Name,
		Project_id:   c.Project_Id,
		CreateDate:   c.CreateDate,
		UpdateDate:   c.UpdateDate,
	}
}

func (c Component) ToAddComponentResponseDto() *dto.AddComponentResponse {
	return &dto.AddComponentResponse{c.Component_Id, c.Name, c.Project_Id, c.CreateDate, c.UpdateDate}
}

type ComponentRepository interface {
	AddComponent(component Component, projectId string) (*Component, *errs.AppError)
	AllComponent(projectKey string, pageId int) ([]Component, *errs.AppError)
	GetComponent(id string) (*Component, *errs.AppError)
	DeleteComponent(id string) *errs.AppError
	UpdateComponent(id string, newComponent Component) *errs.AppError
}
