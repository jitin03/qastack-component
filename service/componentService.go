package service

import (
	"qastack-components/domain"
	"qastack-components/dto"
	"qastack-components/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type ComponentService interface {
	AddComponent(request dto.AddComponentRequest, projectId string) (*dto.AddComponentResponse, *errs.AppError)
	AllComponent(projectKey string, pageId int) ([]dto.AddComponentResponse, *errs.AppError)
	GetComponent(id string) (*dto.AddComponentResponse, *errs.AppError)
	DeleteComponent(id string) (dto.DeleteComponentResponse, *errs.AppError)
	UpdateComponent(id string, request dto.UpdateComponentRequest) (dto.UpdateComponentResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.ComponentRepository
}

func (s DefaultUserService) DeleteComponent(id string) (dto.DeleteComponentResponse, *errs.AppError) {

	err := s.repo.DeleteComponent(id)
	if err != nil {
		return dto.DeleteComponentResponse{
			Message: "Not Delete",
		}, err
	}
	message := dto.DeleteComponentResponse{
		Message: "Deleted",
	}
	return message, nil
}

func (s DefaultUserService) UpdateComponent(id string, req dto.UpdateComponentRequest) (dto.UpdateComponentResponse, *errs.AppError) {
	c := domain.Component{
		Name:       req.Name,
		Project_Id: req.Project_id,
		UpdateDate: time.Now().Format(dbTSLayout),
	}
	err := s.repo.UpdateComponent(id, c)
	if err != nil {
		return dto.UpdateComponentResponse{
			Message: "Not Updated",
		}, err
	}

	message := dto.UpdateComponentResponse{
		Message: "Component is updated ",
	}
	return message, nil

}
func (s DefaultUserService) AllComponent(projectKey string, pageId int) ([]dto.AddComponentResponse, *errs.AppError) {

	components, err := s.repo.AllComponent(projectKey, pageId)
	if err != nil {
		return nil, err
	}
	response := make([]dto.AddComponentResponse, 0)
	for _, component := range components {
		response = append(response, component.ToDto())
	}
	return response, err
}

func (s DefaultUserService) GetComponent(id string) (*dto.AddComponentResponse, *errs.AppError) {
	component, err := s.repo.GetComponent(id)
	if err != nil {
		return nil, err
	}
	response := component.ToDto()

	return &response, err
}

func (s DefaultUserService) AddComponent(req dto.AddComponentRequest, projectId string) (*dto.AddComponentResponse, *errs.AppError) {

	c := domain.Component{

		Name:       req.Name,
		Project_Id: req.Project_id,
		CreateDate: time.Now().Format(dbTSLayout),
		UpdateDate: time.Now().Format(dbTSLayout),
	}

	if newComponent, err := s.repo.AddComponent(c, projectId); err != nil {
		return nil, err
	} else {
		return newComponent.ToAddComponentResponseDto(), nil
	}
}

func NewComponentService(repository domain.ComponentRepository) DefaultUserService {
	return DefaultUserService{repository}
}
